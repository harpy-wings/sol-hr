package models

import (
	"errors"
	"time"
)

type MLeave struct {
	ID             int64          `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	SolderUid      string         `json:"solder_uid" gorm:"not null;varchar(11)" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b"`
	LeaveRequestId int64          `json:"leave_request_id" gorm:"not null" example:"1"`
	LeaveRequest   *MLeaveRequest `json:"leave_request" gorm:"foreignKey:LeaveRequestId"`
	Date           time.Time      `json:"date" gorm:"not null" example:"2021-01-01"`
	Kind           int            `json:"kind" gorm:"not null" example:"1"`
	CreatedAt      time.Time      `json:"created_at" gorm:"autoCreateTime" example:"2021-01-01 00:00:00"`
	UpdatedAt      time.Time      `json:"updated_at" gorm:"autoUpdateTime" example:"2021-01-01 00:00:00"`
}

func (m *MLeave) TableName() string {
	return "leaves"
}

func (m *MLeave) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MLeave{})
	if err != nil {
		return err
	}
	return nil
}
