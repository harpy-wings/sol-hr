package models

import "errors"

type MMilateryBase struct {
	ID                            int64      `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	Title                         string     `json:"title" gorm:"type:varchar(255);not null" example:"قرارگاه"`
	LocationID                    int64      `json:"location_id" gorm:"not null" example:"1"`
	Location                      *MLocation `json:"location" gorm:"foreignKey:LocationID"`
	NativeServiceDurationMonth    int        `json:"native_service_duration_month" gorm:"not null" example:"1"`
	NonNativeServiceDurationMonth int        `json:"non_native_service_duration_month" gorm:"not null" example:"1"`
	Alias                         string     `json:"alias" gorm:"type:varchar(255);not null" example:"qrgg"`
}

func (m *MMilateryBase) TableName() string {
	return "milatery_bases"
}

func (m *MMilateryBase) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MMilateryBase{})
	if err != nil {
		return err
	}

	return nil
}
