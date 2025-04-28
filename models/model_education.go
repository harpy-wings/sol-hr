package models

import "errors"

type MEducationLevel struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"لیسانس"`
	Alias string `json:"alias" gorm:"type:varchar(255);not null" example:"bachelor"`
}

func (m *MEducationLevel) TableName() string {
	return "education_level"
}

func (m *MEducationLevel) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MEducationLevel{})
	if err != nil {
		return err
	}
	educationLevels := []MEducationLevel{}
	err = DB.Find(&educationLevels).Error
	if err != nil {
		return err
	}
	if len(educationLevels) > 0 {
		return nil
	}
	educationLevels = []MEducationLevel{

		{ID: 1, Title: "سیکل و کمتر", Alias: "cycle"},
		{ID: 2, Title: "دیپلم", Alias: "diploma"},
		{ID: 3, Title: "فوق دیپلم", Alias: "کاردانی"},
		{ID: 4, Title: "لیسانس", Alias: "کارشناسی"},
		{ID: 5, Title: "فوق لیسانس", Alias: "کارشناسی ارشد"},
		{ID: 6, Title: "دکترا", Alias: "دکتری تخصصی"},
	}
	err = DB.CreateInBatches(educationLevels, 100).Error
	if err != nil {
		return err
	}
	return nil
}
