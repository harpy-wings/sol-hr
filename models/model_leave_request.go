package models

import (
	"errors"
	"time"
)

type MLeaveRequest struct {
	ID        int64     `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	SolderUid string    `json:"solder_uid" gorm:"not null;varchar(11)" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b"`
	StartDate time.Time `json:"start_date" gorm:"not null" example:"2021-01-01 00:00:00"`
	DueDate   time.Time `json:"due_date" gorm:"not null" example:"2021-01-01 00:00:00"`

	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" example:"2021-01-01 00:00:00"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" example:"2021-01-01 00:00:00"`
	Status    int       `json:"status" gorm:"not null;default:0" example:"0"`

	AcceptedById int64  `json:"accepted_by_id" gorm:"not null" example:"1"`
	AcceptedBy   *MUser `json:"accepted_by" gorm:"foreignKey:AcceptedById"`

	RequestedById int64  `json:"requested_by_id" gorm:"not null" example:"1"`
	RequestedBy   *MUser `json:"requested_by" gorm:"foreignKey:RequestedById"`

	Reason string `json:"reason" gorm:"" example:"شخصی"`
	Kind   int    `json:"kind" gorm:"not null;default:0" example:"0"`
}

type MLeaveRequestStatus int

const (
	MLeaveRequestStatusPending MLeaveRequestStatus = iota
	MLeaveRequestStatusAccepted
	MLeaveRequestStatusRejected
)

type MLeaveRequestKind int

const (
	MLeaveRequestKindAnnual MLeaveRequestKind = iota
	MLeaveRequestKindHolodayBonus
	MLeaveRequestKindBereavement
	MLeaveRequestKindReward
	MLeaveRequestKindSick
)

func (m *MLeaveRequest) TableName() string {
	return "leave_requests"
}

func (m *MLeaveRequest) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MLeaveRequest{})
	if err != nil {
		return err
	}
	return nil
}
