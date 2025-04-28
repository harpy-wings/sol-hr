package models

import (
	"errors"
	"time"
)

type MUserSession struct {
	UserSessionID int64      `json:"user_session_id" gorm:"primary_key;auto_increment"`
	UserID        int64      `json:"user_id" gorm:"not null"`
	Secret        string     `json:"secret" gorm:"not null"`
	ExpiresAt     time.Time  `json:"expires_at" gorm:"not null"`
	LogoutAt      *time.Time `json:"logout_at"`
	CreatedAt     time.Time  `json:"created_at" gorm:"not null"`
	UpdatedAt     time.Time  `json:"updated_at" gorm:"not null"`
	Device        string     `json:"device" gorm:"not null"`
	IP            string     `json:"ip" gorm:"not null"`
	IsSA          bool       `json:"is_sa" gorm:"not null"`
}

func (*MUserSession) TableName() string {
	return "user_session"
}

func (m *MUserSession) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MUserSession{})
	if err != nil {
		return err
	}

	return nil
}
