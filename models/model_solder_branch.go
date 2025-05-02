package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type MSolderBranch struct {
	gorm.Model
	SolderUid string   `json:"solder_uid" gorm:"not null" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b"`
	Solder    *MSolder `json:"solder" gorm:"-"`

	BranchID int64    `json:"branch_id" gorm:"not null" example:"1"`
	Branch   *MBranch `json:"branch" gorm:"-"`

	StartDate time.Time  `json:"start_date" gorm:"not null" example:"2021-01-01"`
	DueDate   *time.Time `json:"due_date" gorm:"" example:"2021-01-01"`

	IsActive bool `json:"is_active" gorm:"not null;default:true" example:"true"`
}

func (m *MSolderBranch) TableName() string {
	return "solder_branch"
}

func (m *MSolderBranch) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MSolderBranch{})
	if err != nil {
		return err
	}
	return nil
}
