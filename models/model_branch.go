package models

import (
	"errors"
	"time"
)

type MBranch struct {
	ID                            int64      `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	Xid                           int64      `json:"xid" gorm:"not null" example:"1"` // xid 0 is internal unit, 1 is external milatery base
	Title                         string     `json:"title" gorm:"type:varchar(255);not null" example:"قرارگاه"`
	ParrentID                     int64      `json:"parrent_id" gorm:"not null" example:"1"`
	Parrent                       *MBranch   `json:"-" gorm:"-"`
	Children                      []*MBranch `json:"children" gorm:"-"`
	Alias                         string     `json:"alias" gorm:"type:varchar(255);not null" example:"qrgg"`
	LocationID                    int64      `json:"location_id" gorm:"not null" example:"1"`
	Location                      *MLocation `json:"location" gorm:"-"`
	NativeServiceDurationMonth    int        `json:"native_service_duration_month" gorm:"not null" example:"1"`
	NonNativeServiceDurationMonth int        `json:"non_native_service_duration_month" gorm:"not null" example:"1"`
	ExpectedSolderGradedCount     int        `json:"expected_solder_graded_count" gorm:"not null;default:0" example:"1"`
	ActualSolderGradedCount       int        `json:"actual_solder_graded_count" gorm:"not null;default:0" example:"1"`
	ExpectedSolderNonGradedCount  int        `json:"expected_solder_non_graded_count" gorm:"not null;default:0" example:"1"`
	ActualSolderNonGradedCount    int        `json:"actual_solder_non_graded_count" gorm:"not null;default:0" example:"1"`
	CreatedAt                     time.Time  `json:"created_at" gorm:"not null" example:"2021-01-01T00:00:00Z"`
	UpdatedAt                     time.Time  `json:"updated_at" gorm:"not null" example:"2021-01-01T00:00:00Z"`
}

func (m *MBranch) TableName() string {
	return "branch"
}

func (m *MBranch) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MBranch{})
	if err != nil {
		return err
	}
	branches := []MBranch{}
	err = DB.Find(&branches).Error
	if err != nil {
		return err
	}
	if len(branches) > 0 {
		return nil
	}
	branches = []MBranch{

		{ID: 1, Title: "قرارگاه", ParrentID: 0, Alias: "", Xid: 0},
		{ID: 2, Title: "سایت 2", ParrentID: 1, Alias: "", Xid: 0},
		{ID: 3, Title: "سایت 3", ParrentID: 1, Alias: "", Xid: 0},
		{ID: 4, Title: "کشت و صنعت", ParrentID: 0, Alias: "", Xid: 0},
		{ID: 5, Title: "نیرو انسانی", ParrentID: 0, Alias: "", Xid: 0},
		{ID: 6, Title: "گروه عملیات شهرآباد", ParrentID: 0, Alias: "", Xid: 1, LocationID: 1, NativeServiceDurationMonth: 18, NonNativeServiceDurationMonth: 15},
	}
	err = DB.CreateInBatches(branches, 100).Error
	if err != nil {
		return err
	}
	return nil
}
