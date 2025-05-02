package models

import "errors"

type MSolderLeaveProfile struct {
	ID        int64    `json:"id" gorm:"primaryKey;autoIncrement:true" example:"1"`
	SolderUid string   `json:"solder_uid" gorm:"type:varchar(11);not null" example:"1234567890"`
	Solder    *MSolder `json:"solder" gorm:"-"`

	TotalAnnualLeave int64 `json:"total_annual_leave" gorm:"not null" example:"10"`
	UsedAnnualLeave  int64 `json:"used_annual_leave" gorm:"not null" example:"10"`

	TotalHolidayBonus int64 `json:"total_holiday_bonus" gorm:"not null" example:"10"`
	UsedHolidayBonus  int64 `json:"used_holiday_bonus" gorm:"not null" example:"10"`

	TotalBereavementLeave int64 `json:"total_bereavement_leave" gorm:"not null" example:"10"`
	UsedBereavementLeave  int64 `json:"used_bereavement_leave" gorm:"not null" example:"10"`

	TotalRewardLeave int64 `json:"total_reward_leave" gorm:"not null" example:"10"`
	UsedRewardLeave  int64 `json:"used_reward_leave" gorm:"not null" example:"10"`

	TotalSickLeave int64 `json:"total_sick_leave" gorm:"not null" example:"10"`
	UsedSickLeave  int64 `json:"used_sick_leave" gorm:"not null" example:"10"`

	TotalNonNativeTravelLeave int64 `json:"total_non_native_travel_leave" gorm:"not null" example:"10"`
	UsedNonNativeTravelLeave  int64 `json:"used_non_native_travel_leave" gorm:"not null" example:"10"`

	TotalOtherLeave int64 `json:"total_other_leave" gorm:"not null" example:"10"`
	UsedOtherLeave  int64 `json:"used_other_leave" gorm:"not null" example:"10"`
}

func (m *MSolderLeaveProfile) TableName() string {
	return "solder_leave_profiles"
}

func (m *MSolderLeaveProfile) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MSolderLeaveProfile{})
	if err != nil {
		return err
	}
	return nil
}
