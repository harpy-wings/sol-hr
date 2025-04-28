package models

import "errors"

type MReligon struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"مسیحی"`
}

func (m *MReligon) TableName() string {
	return "religons"
}

func (m *MReligon) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MReligon{})
	if err != nil {
		return err
	}
	religons := []MReligon{}
	err = DB.Find(&religons).Error
	if err != nil {
		return err
	}
	if len(religons) > 0 {
		return nil
	}
	religons = []MReligon{
		{ID: 1, Title: "شیعه"},
		{ID: 2, Title: "سنی"},
		{ID: 3, Title: "مسیحی"},
		{ID: 4, Title: "یهودی"},
		{ID: 5, Title: "زرتشتی"},
		{ID: 6, Title: "سایر"},
	}
	err = DB.CreateInBatches(religons, 100).Error
	if err != nil {
		return err
	}
	return nil
}
