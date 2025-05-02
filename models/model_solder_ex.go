package models

import (
	"errors"
	"time"

	"github.com/lib/pq"
)

type MExSolder struct {
	Uid string `json:"uid" gorm:"primary_key;not null" example:"1dec3bcb-7b1a-4b90-b261-6763da09e06b"`

	MariageDate   *time.Time `json:"mariage_date" gorm:"" example:"2021-01-01"`
	IsMaried      bool       `json:"is_maried" gorm:"not null;default:false" example:"false"`
	ChildrenCount int64      `json:"children_count" gorm:"not null;default:0" example:"0"`

	ReligionID            int64          `json:"religion_id" gorm:"not null" example:"1"`
	Religion              *MReligon      `json:"religion" gorm:"-"`
	BirthDate             *time.Time     `json:"birth_date" gorm:"" example:"2021-01-01"`
	BirthPlace            string         `json:"birth_place" gorm:"" example:"تهران"`
	BirthIssuedPlace      string         `json:"birth_issued_place" gorm:"" example:"تهران"`
	BankAccount           string         `json:"bank_account" gorm:"" example:"1234567890"`
	HightCm               int64          `json:"hight_cm" gorm:"not null" example:"170"`
	WeightKg              int64          `json:"weight_kg" gorm:"not null" example:"70"`
	BloodType             string         `json:"blood_type" gorm:"" example:"A+"`
	Skills                pq.StringArray `json:"skills" gorm:"type:text[]"`
	FatherPhone           string         `json:"father_phone" gorm:"" example:"09123456789"`
	ExtraPhoneNumber      string         `json:"extra_phone_number" gorm:"" example:"09123456789"`
	DriverLicenseIssuedAt *time.Time     `json:"driver_license_issued_at" gorm:"" example:"2021-01-01"`

	EyeColorID  int64   `json:"eye_color_id" gorm:"not null" example:"1"`
	EyeColor    *MColor `json:"eye_color" gorm:"-"`
	HairColorID int64   `json:"hair_color_id" gorm:"not null" example:"1"`
	HairColor   *MColor `json:"hair_color" gorm:"-"`

	EducationField      string           `json:"education_field" gorm:"" example:"مهندسی کامپیوتر"`
	EducationUniversity string           `json:"education_university" gorm:"" example:"تهران"`
	EducationLevelID    int64            `json:"education_level_id" gorm:"not null" example:"1"`
	EducationLevel      *MEducationLevel `json:"education_level" gorm:"-"`
}

func (m *MExSolder) TableName() string {
	return "solder_ex"
}

func (m *MExSolder) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MExSolder{})
	if err != nil {
		return err
	}
	return nil
}
