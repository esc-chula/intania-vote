package model

import "github.com/google/uuid"

type User struct {
	Model

	Ballots []Ballot `gorm:"foreignKey:UserId"`
	Votes   []Vote   `gorm:"foreignKey:UserId"`

	OidcId    uuid.UUID `gorm:"size:36;uniqueIndex"`
	StudentId string    `gorm:"size:36;not null;uniqueIndex"`
}

func (User) TableName() string {
	return "users"
}
