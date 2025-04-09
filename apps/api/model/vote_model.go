package model

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Owner string

const (
	Student Owner = "STUDENT"
	ESC     Owner = "ESC"
	ISESC   Owner = "ISESC"
)

func (s *Owner) Scan(value interface{}) error {
	strVal, ok := value.(string)
	if !ok {
		return fmt.Errorf("invalid data type for Owner: %T", value)
	}

	switch Owner(strVal) {
	case Student, ESC, ISESC:
		*s = Owner(strVal)
		return nil
	default:
		return fmt.Errorf("invalid value for Owner: %s", strVal)
	}
}

func (s Owner) Value() (driver.Value, error) {
	switch s {
	case Student, ESC, ISESC:
		return string(s), nil
	default:
		return nil, fmt.Errorf("invalid value for Owner: %s", s)
	}
}

type Vote struct {
	Model

	UserId  uint     `gorm:"not null"`
	User    User     `gorm:"foreignKey:UserId"`
	Choices []Choice `gorm:"foreignKey:VoteId"`

	Name               string    `gorm:"size:100;not null"`
	Description        string    `gorm:"size:300"`
	Slug               string    `gorm:"size:100;not null;uniqueIndex"`
	Image              *string   `gorm:"size:100"`
	Owner              Owner     `gorm:"type:owner;not null"`
	EligibleStudentId  string    `gorm:"size:50"`
	EligibleDepartment string    `gorm:"size:50"`
	EligibleYear       string    `gorm:"size:50"`
	IsPrivate          bool      `gorm:"not null;default:false"`
	IsRealTime         bool      `gorm:"not null;default:false"`
	IsAllowMultiple    bool      `gorm:"not null;default:false"`
	StartAt            time.Time `gorm:"not null"`
	EndAt              time.Time `gorm:"not null"`
}

func (Vote) TableName() string {
	return "votes"
}
