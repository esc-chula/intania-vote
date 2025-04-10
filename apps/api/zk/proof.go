package zk

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/big"
)

type Proof struct {
	Commitment     string `json:"commitment"`
	BlindingFactor string `json:"blinding_factor"`
	Challenge      string `json:"challenge"`
	Response       string `json:"response"`
	Nullifier      string `json:"nullifier"`
	Receipt        string `json:"receipt"`
}

func GenerateProof(choiceId uint, secret string) (*Proof, error) {
	// Convert choiceId to big.Int
	m := new(big.Int).SetUint64(uint64(choiceId))

	// Generate random blinding factor
	r, err := GenerateRandomBigInt(PrimeModulus)
	if err != nil {
		return nil, err
	}

	// Create commitment
	commitment := CreatePedersenCommitment(m, r)

	// Create a nullifier (unique identifier for this vote)
	nullifierBytes := make([]byte, 32)
	_, err = rand.Read(nullifierBytes)
	if err != nil {
		return nil, err
	}
	nullifier := hex.EncodeToString(nullifierBytes)

	// Generate a challenge - in a real system, this would involve more complex protocol
	secretBytes := []byte(secret)
	challengeInput := append(commitment.Bytes(), secretBytes...)
	challengeInput = append(challengeInput, nullifierBytes...)
	challengeInput = append(challengeInput, m.Bytes()...) // Include candidate ID in challenge
	challengeHash := sha256.Sum256(challengeInput)
	challenge := new(big.Int).SetBytes(challengeHash[:])

	// Generate response: s = r + e*m mod (p-1)
	// Where e is the challenge, r is the blinding factor, m is the vote
	pMinus1 := new(big.Int).Sub(PrimeModulus, big.NewInt(1))
	em := new(big.Int).Mul(challenge, m)
	em.Mod(em, pMinus1)

	response := new(big.Int).Add(r, em)
	response.Mod(response, pMinus1)

	// Generate receipt
	// timestamp := time.Now().Unix()
	// NEW: Generate receipt only returned to user, not stored
	receiptData := fmt.Sprintf("%s:%s:%d", commitment.String(), nullifier, choiceId)
	receiptHash := sha256.Sum256([]byte(receiptData))
	receipt := base64.StdEncoding.EncodeToString(receiptHash[:])

	// Encrypt the blinding factor with the secret
	// In a real system, this would use proper encryption
	blindingFactorEncrypted := EncryptBlindingFactor(r, secret)

	// Build proof object
	proofResponse := &Proof{
		Commitment:     commitment.String(),
		BlindingFactor: blindingFactorEncrypted,
		Challenge:      challenge.String(),
		Response:       response.String(),
		Nullifier:      nullifier,
		Receipt:        receipt,
	}

	return proofResponse, nil
}

func VerifyProof(proofData map[string]string, choiceId uint, secret string) bool {
	// Basic validation - check that all required fields exist
	requiredFields := []string{"commitment", "challenge", "response", "nullifier", "blindingFactor"}
	for _, field := range requiredFields {
		if _, exists := proofData[field]; !exists {
			return false
		}
	}

	// Parse the commitment
	commitment, ok := new(big.Int).SetString(proofData["commitment"], 10)
	if !ok {
		return false
	}

	// Parse the challenge
	_, ok = new(big.Int).SetString(proofData["challenge"], 10)
	if !ok {
		return false
	}

	// Parse the response
	_, ok = new(big.Int).SetString(proofData["response"], 10)
	if !ok {
		return false
	}

	// Decrypt the blinding factor
	blindingFactor, err := DecryptBlindingFactor(proofData["blindingFactor"], secret)
	if err != nil {
		return false
	}

	// Convert choiceId to big.Int
	m := new(big.Int).SetUint64(uint64(choiceId))

	// Verify the commitment: Check if commitment = g^m * h^r
	expectedCommitment := CreatePedersenCommitment(m, blindingFactor)
	if commitment.Cmp(expectedCommitment) != 0 {
		return false
	}

	// In a real system, we would also verify the Schnorr proof
	// For this simplified example, we just check that the commitment matches

	return true
}
