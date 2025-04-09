package model

type Ballot struct {
	Model

	UserId uint `gorm:"not null"`
	User   User `gorm:"foreignKey:UserId"`
	VoteId uint `gorm:"not null"`
	Vote   Vote `gorm:"foreignKey:VoteId"`

	Challenge      string `gorm:"not null"`
	Commitment     string `gorm:"not null"`
	BlindingFactor string `gorm:"not null"`
	Response       string `gorm:"not null"`
	Encrypted      string `gorm:"not null"`
	Nullifier      string `gorm:"uniqueIndex;not null"`
}

func (Ballot) TableName() string {
	return "ballots"
}
