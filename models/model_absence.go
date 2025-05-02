package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type MAbsence struct {
	gorm.Model
	SolderUid   string    `json:"solder_uid" gorm:"type:varchar(11);not null" example:"1234567890"`
	Solder      *MSolder  `json:"solder" gorm:"-"`
	Date        time.Time `json:"date" gorm:"type:date;not null" example:"2021-01-01"`
	CreatedByID int64     `json:"created_by_id" gorm:"not null" example:"1"`
	CreatedBy   *MUser    `json:"created_by" gorm:"-"`

	IsProcessed  bool         `json:"is_processed" gorm:"not null" example:"false"`
	AdjustmentID *int64       `json:"adjustment_id" gorm:"" example:"1"`
	Adjustment   *MAdjustment `json:"adjustment" gorm:"-"`
}

func (m *MAbsence) TableName() string {
	return "absences"
}

func (m *MAbsence) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}

	err := DB.AutoMigrate(&MAbsence{})
	if err != nil {
		return err
	}

	return nil
}
