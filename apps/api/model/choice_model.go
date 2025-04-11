package model

type Choice struct {
	Model

	VoteId uint `gorm:"not null"`
	Vote   Vote `gorm:"foreignKey:VoteId"`

	Number        uint32  `gorm:"size:10;not null"`
	Name          string  `gorm:"size:50;not null"`
	Description   string  `gorm:"size:50"`
	Information   *string `gorm:"type:text"`
	Image         *string `gorm:"size:100"`
	BallotCounter uint32  `gorm:"default:0"`
}

func (Choice) TableName() string {
	return "choices"
}
