package models

import (
	"errors"
	"time"
)

type MSolder struct {
	Uid         string `json:"uid" gorm:"primary_key;not null" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b"`
	Personel_id string `json:"personel_id" gorm:"not null" example:"1234567890"`
	FirstName   string `json:"first_name" gorm:"not null" example:"John"`
	LastName    string `json:"last_name" gorm:"not null" example:"Doe"`
	FatherName  string `json:"father_name" gorm:"not null" example:"John"`

	PrimaryBranchID int64    `json:"primary_branch_id" gorm:"not null" example:"1"`
	PrimaryBranch   *MBranch `json:"primary_branch" gorm:"foreignKey:PrimaryBranchID"`

	SecondaryBranchID int64    `json:"secondary_branch_id" gorm:"not null" example:"1"`
	SecondaryBranch   *MBranch `json:"secondary_branch" gorm:"foreignKey:SecondaryBranchID"`

	MilitaryRankID int64          `json:"military_rank_id" gorm:"not null" example:"1"`
	MilitaryRank   *MMilitaryRank `json:"military_rank" gorm:"foreignKey:MilitaryRankID"`

	DischargeNumber int64     `json:"discharge_number" gorm:"not null" example:"1234567890"`
	DischargeDate   time.Time `json:"discharge_date" gorm:"not null" example:"2021-01-01"`

	HasDisability     bool   `json:"has_disability" gorm:"not null;default:false" example:"false"`
	DisabilityComment string `json:"disability_comment" gorm:"" example:"1"`

	IsMentallyHealthy      bool   `json:"is_mentally_healthy" gorm:"not null;default:false" example:"false"`
	MentallyHealthyComment string `json:"mentally_healthy_comment" gorm:"" example:"1"`

	LeaveProfileID int64                `json:"leave_profile_id" gorm:"not null" example:"1"`
	LeaveProfile   *MSolderLeaveProfile `json:"leave_profile" gorm:"foreignKey:LeaveProfileID"`
}

func (m *MSolder) TableName() string {
	return "solder"
}

func (m *MSolder) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MSolder{})
	if err != nil {
		return err
	}
	return nil
}
