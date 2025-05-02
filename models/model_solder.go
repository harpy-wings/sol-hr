package models

import (
	"errors"
	"time"
)

type MSolder struct {
	Uid         string `json:"uid" gorm:"primary_key;not null" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b" validate:"required,len=10"`
	Personel_id string `json:"personel_id" gorm:"not null" example:"1234567890" validate:"required,len=9"`
	FirstName   string `json:"first_name" gorm:"not null" example:"John" validate:"required"`
	LastName    string `json:"last_name" gorm:"not null" example:"Doe" validate:"required"`
	FatherName  string `json:"father_name" gorm:"not null" example:"John" validate:"required"`

	PrimaryBranchID int64    `json:"primary_branch_id" gorm:"not null" example:"1"`
	PrimaryBranch   *MBranch `json:"primary_branch" gorm:"-"`

	SecondaryBranchID int64    `json:"secondary_branch_id" gorm:"not null" example:"1"`
	SecondaryBranch   *MBranch `json:"secondary_branch" gorm:"-"`

	MilitaryRankID int64          `json:"military_rank_id" gorm:"not null" example:"1"`
	MilitaryRank   *MMilitaryRank `json:"military_rank" gorm:"-"`

	SeviceStartedAt      time.Time `json:"sevice_started_at" gorm:"" example:"2021-01-01"`
	ServiceExpectedEndAt time.Time `json:"service_expected_end_at" gorm:"" example:"2021-01-01"`
	DischargeNumber      int64     `json:"discharge_number" gorm:"not null" example:"1234567890"`
	DischargeDate        time.Time `json:"discharge_date" gorm:"not null" example:"2021-01-01"`

	HasDisability     bool   `json:"has_disability" gorm:"not null;default:false" example:"false"`
	DisabilityComment string `json:"disability_comment" gorm:"" example:"1"`

	IsMentallyHealthy      bool   `json:"is_mentally_healthy" gorm:"not null;default:false" example:"false"`
	MentallyHealthyComment string `json:"mentally_healthy_comment" gorm:"" example:"1"`

	LeaveProfileID int64                `json:"leave_profile_id" gorm:"not null" example:"1"`
	LeaveProfile   *MSolderLeaveProfile `json:"leave_profile" gorm:"-"`
	Ditails        *MExSolder           `json:"ditails" gorm:"-"`

	Status MSolderStatus `json:"status" gorm:"not null;default:1" example:"1"`

	CreatedAt time.Time `json:"created_at" gorm:"not null;default:CURRENT_TIMESTAMP" example:"2021-01-01"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null;default:CURRENT_TIMESTAMP" example:"2021-01-01"`

	CreatedByID int64  `json:"created_by_id" gorm:"not null" example:"1"`
	CreatedBy   *MUser `json:"created_by" gorm:"-"`
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

type MSolderStatus int64

const (
	MSolderStatusUnknown                MSolderStatus = iota
	MSolderStatusAbsent                               // غیبت
	MSolderStatusEscaped                              // فرار
	MSolderStatusEscapedMoreThan6Months               // فرار بیش از شش ماه
	MSolserStatusServing                              // درحال خدمت
	MSolderStatusLeave                                // مرخصی
	MSolderStatusExempt                               // معافیت
	MSolderStatusRetired                              // ترخیص
	// MSolderStatusDeceased                             // فوت
)
