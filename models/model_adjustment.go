package models

import (
	"errors"

	"gorm.io/gorm"
)

type MAdjustment struct {
	gorm.Model
	SolderUid   string   `json:"solder_uid" gorm:"type:varchar(11);not null" example:"1234567890"`
	Solder      *MSolder `json:"solder" gorm:"-"`
	Title       string   `json:"title" gorm:"type:varchar(255);not null" example:"20 days extra"`
	DurationDay int64    `json:"duration_day" gorm:"not null" example:"20"`

	CategoryID int64                `json:"category_id" gorm:"not null" example:"1"`
	Category   *MAdjustmentCategory `json:"category" gorm:"-"`

	CreatedByID int64  `json:"created_by_id" gorm:"not null" example:"1"`
	CreatedBy   *MUser `json:"created_by" gorm:"-"`
}

func (m *MAdjustment) TableName() string {
	return "adjustments"
}

func (m *MAdjustment) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MAdjustment{})
	if err != nil {
		return err
	}
	return nil
}
