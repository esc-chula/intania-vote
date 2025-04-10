package zk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/esc-chula/intania-vote/apps/api/config"
)

var (
	// Prime modulus (large safe prime)
	PrimeModulus, _ = new(big.Int).SetString("FFFFFFFFFFFFFFFFC90FDAA22168C234C4C6628B80DC1CD129024E088A67CC74020BBEA63B139B22514A08798E3404DDEF9519B3CD3A431B302B0A6DF25F14374FE1356D6D51C245E485B576625E7EC6F44C42E9A637ED6B0BFF5CB6F406B7EDEE386BFB5A899FA5AE9F24117C4B1FE649286651ECE45B3DC2007CB8A163BF0598DA48361C55D39A69163FA8FD24CF5F83655D23DCA3AD961C62F356208552BB9ED529077096966D670C354E4ABC9804F1746C08CA18217C32905E462E36CE3BE39E772C180E86039B2783A2EC07A28FB5C55DF06F4C52C9DE2BCBF6955817183995497CEA956AE515D2261898FA051015728E5A8AACAA68FFFFFFFFFFFFFFFF", 16)

	// Generator for Pedersen commitment
	Generator = big.NewInt(2)

	// Second generator for Pedersen commitment
	H, _ = new(big.Int).SetString("5C7FF6B06F8F143FE8288433493E4769C4D988ACE5BE25A0E24809670716C613D7B0CEE6932F8FAA7C44D2CB24523DA53FBE4F6EC3595892D1AA58C4328A06C46A15662E7EAA703A1DECF8BBB2D05DBE2EB956C142A338661D10461C0D135472085057F3494309FFA73C611F78B32ADBB5740C361C9F35BE90997DB2014E2EF5", 16)
)

// generateRandomBigInt generates a random big integer within the range [1, max-1]
func GenerateRandomBigInt(max *big.Int) (*big.Int, error) {
	// Generate random bytes
	b := make([]byte, max.BitLen()/8+1)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	// Convert to big int and ensure it's in range
	r := new(big.Int).SetBytes(b)
	r.Mod(r, max)

	// Ensure it's not zero
	if r.Cmp(big.NewInt(0)) == 0 {
		return GenerateRandomBigInt(max)
	}

	return r, nil
}

// createPedersenCommitment creates a Pedersen commitment to a value
// Commitment = g^m * h^r mod p
func CreatePedersenCommitment(value *big.Int, blindingFactor *big.Int) *big.Int {
	// g^m mod p
	gm := new(big.Int).Exp(Generator, value, PrimeModulus)

	// h^r mod p
	hr := new(big.Int).Exp(H, blindingFactor, PrimeModulus)

	// g^m * h^r mod p
	commitment := new(big.Int).Mul(gm, hr)
	commitment.Mod(commitment, PrimeModulus)

	return commitment
}

// Simple encryption of blinding factor using user's secret
// In a real system, this would use proper encryption
func EncryptBlindingFactor(r *big.Int, secret string) string {
	secretHash := sha256.Sum256([]byte(secret))
	secretInt := new(big.Int).SetBytes(secretHash[:])

	// Simple XOR-like encryption (for demonstration only)
	// In a real system, use proper encryption
	encrypted := new(big.Int).Xor(r, secretInt)
	return encrypted.String()
}

// Decrypt the blinding factor using user's secret
func DecryptBlindingFactor(encrypted string, secret string) (*big.Int, error) {
	encryptedInt, ok := new(big.Int).SetString(encrypted, 10)
	if !ok {
		return nil, fmt.Errorf("invalid encrypted blinding factor")
	}

	secretHash := sha256.Sum256([]byte(secret))
	secretInt := new(big.Int).SetBytes(secretHash[:])

	// Reverse the XOR operation
	decrypted := new(big.Int).Xor(encryptedInt, secretInt)
	return decrypted, nil
}

// Helper function to encrypt with server key
func EncryptWithServerKey(choiceId int, cfg *config.Config) string {
	data := []byte(fmt.Sprintf("%d", choiceId))

	h := hmac.New(sha256.New, []byte(cfg.HmacKey))
	h.Write([]byte(cfg.HmacSalt))
	key := h.Sum(nil)

	// If key is too short for AES, extend it
	if len(key) < 32 {
		h := sha256.Sum256(key)
		key = h[:]
	}

	// Create cipher
	block, err := aes.NewCipher(key[:32]) // Use first 32 bytes for AES-256
	if err != nil {
		// Fallback to simple XOR
		return base64.StdEncoding.EncodeToString(XorEncrypt(data, key))
	}

	// Generate IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return base64.StdEncoding.EncodeToString(XorEncrypt(data, key))
	}

	// Encrypt
	ciphertext := make([]byte, len(data))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, data)

	// Combine IV and ciphertext
	result := append(iv, ciphertext...)
	return base64.StdEncoding.EncodeToString(result)
}

// Helper function to decrypt with server key
func DecryptWithServerKey(encrypted string, cfg *config.Config) (int, error) {
	h := hmac.New(sha256.New, []byte(cfg.HmacKey))
	h.Write([]byte(cfg.HmacSalt))
	key := h.Sum(nil)

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return 0, err
	}

	// If key is too short for AES, extend it
	if len(key) < 32 {
		h := sha256.Sum256(key)
		key = h[:]
	}

	// Try AES decryption
	if len(ciphertext) > aes.BlockSize {
		// Extract IV
		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]

		// Create cipher
		block, err := aes.NewCipher(key[:32])
		if err == nil {
			// Decrypt
			plaintext := make([]byte, len(ciphertext))
			stream := cipher.NewCFBDecrypter(block, iv)
			stream.XORKeyStream(plaintext, ciphertext)

			// Parse
			var choiceId int
			if _, err := fmt.Sscanf(string(plaintext), "%d", &choiceId); err == nil {
				return choiceId, nil
			}
		}
	}

	// Fallback to XOR
	plaintext := XorEncrypt(ciphertext, key)
	var choiceId int
	if _, err := fmt.Sscanf(string(plaintext), "%d", &choiceId); err != nil {
		return 0, err
	}

	return choiceId, nil
}

// Helper function to encrypt candidate ID
func EncryptChoiceId(choiceId int, key []byte) string {

	// Convert candidate ID to bytes
	candidateBytes := []byte(fmt.Sprintf("%d", choiceId))

	// Create cipher with key
	block, err := aes.NewCipher(key)
	if err != nil {
		// Fall back to a simpler encryption if AES fails
		return base64.StdEncoding.EncodeToString(XorEncrypt(candidateBytes, key))
	}

	// Generate random IV
	iv := make([]byte, aes.BlockSize)
	if _, err := rand.Read(iv); err != nil {
		return base64.StdEncoding.EncodeToString(XorEncrypt(candidateBytes, key))
	}

	// Encrypt
	ciphertext := make([]byte, len(candidateBytes))
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext, candidateBytes)

	// Combine IV and ciphertext
	result := append(iv, ciphertext...)

	// log.Println("--- Encrypting with server key ---")
	// log.Println("choiceId:", choiceId)
	// log.Println("key:", string(key))
	// log.Println("encrypted:", base64.StdEncoding.EncodeToString(result))

	return base64.StdEncoding.EncodeToString(result)
}

func DecryptChoiceId(encryptedId string, key []byte) (int, error) {
	// Decode base64
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedId)
	if err != nil {
		return 0, fmt.Errorf("base64 decode error: %w", err)
	}

	// Try AES decryption first
	if len(ciphertext) > aes.BlockSize {
		// Extract IV
		iv := ciphertext[:aes.BlockSize]
		ciphertext = ciphertext[aes.BlockSize:]

		// Create cipher with key
		block, err := aes.NewCipher(key)
		if err == nil {
			// Decrypt
			plaintext := make([]byte, len(ciphertext))
			stream := cipher.NewCFBDecrypter(block, iv)
			stream.XORKeyStream(plaintext, ciphertext)

			// Convert to int
			var choiceId int
			if _, err := fmt.Sscanf(string(plaintext), "%d", &choiceId); err == nil {
				return choiceId, nil
			}

			// If parsing as int failed, but we got valid plaintext
			if len(plaintext) > 0 {
				return 0, fmt.Errorf("decrypted content is not a valid choice ID: %s", string(plaintext))
			}
		}
	}

	// Fallback to XOR decryption
	plaintext := XorEncrypt(ciphertext, key)

	// Try to parse as int
	var choiceId int
	if _, err := fmt.Sscanf(string(plaintext), "%d", &choiceId); err != nil {
		// If it's not a valid int, return a meaningful error
		return 0, fmt.Errorf("decrypted content is not a valid choice ID: %s", string(plaintext))
	}

	// log.Println("--- Decrypting with server key ---")
	// log.Println("encryptedId:", encryptedId)
	// log.Println("key:", string(key))
	// log.Println("choiceId:", choiceId)

	return choiceId, nil
}

// Simple XOR encryption as fallback
func XorEncrypt(data, key []byte) []byte {
	encrypted := make([]byte, len(data))
	for i := 0; i < len(data); i++ {
		encrypted[i] = data[i] ^ key[i%len(key)]
	}
	return encrypted
}

// Generate a unique key for encrypting/decrypting this specific vote
func GenerateVoteEncryptionKey(sub string, nullifier string) []byte {
	// Combine sub and nullifier to generate a unique key
	data := []byte(sub + nullifier)
	hash := sha256.Sum256(data)
	return hash[:] // Return 32-byte key suitable for AES-256
}

// Used internally instead of user-provided secret
func DeriveSecretFromSub(sub string, cfg *config.Config) string {
	key := []byte(cfg.HmacKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(sub))
	return hex.EncodeToString(h.Sum(nil))
}
