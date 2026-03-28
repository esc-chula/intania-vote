package model



type User struct {
	Model

	Ballots []Ballot `gorm:"foreignKey:UserId"`
	Votes   []Vote   `gorm:"foreignKey:UserId"`

	OidcId    string `gorm:"size:255;uniqueIndex"`
	StudentId string `gorm:"size:36;not null;uniqueIndex"`
}

func (User) TableName() string {
	return "users"
}
