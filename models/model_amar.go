package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type MAmar struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	SolderUid string   `json:"solder_uid" gorm:"not null;varchar(11)" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b"`
	Solder    *MSolder `json:"solder" gorm:"-"`

	BranchID int64    `json:"branch_id" gorm:"not null" example:"1"`
	Branch   *MBranch `json:"branch" gorm:"-"`

	Status int `json:"status" gorm:"not null;default:0;int2" example:"0"`

	UserId string `json:"user_id" gorm:"not null;varchar(11)" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b"`
	User   *MUser `json:"user" gorm:"-"`

	IsLast bool `json:"is_last" gorm:"not null;default:false" example:"false"`
}

func (m *MAmar) TableName() string {
	return "amar"
}

func (m *MAmar) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MAmar{})
	if err != nil {
		return err
	}
	return nil
}
