package models

import "errors"

type MColor struct {
	ID    int64  `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"سفید"`
	RGB   string `json:"rgb" gorm:"type:varchar(255);not null" example:"#ffffff"`
}

func (m *MColor) TableName() string {
	return "colors"
}

func (m *MColor) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MColor{})
	if err != nil {
		return err
	}
	colors := []MColor{}
	err = DB.Find(&colors).Error
	if err != nil {
		return err
	}
	if len(colors) > 0 {
		return nil
	}
	colors = []MColor{
		{ID: 1, Title: "سفید", RGB: "#ffffff"},
		{ID: 2, Title: "سیاه", RGB: "#000000"},
		{ID: 3, Title: "قهوه ای", RGB: "#a52a2a"},
		{ID: 4, Title: "آبی", RGB: "#0000ff"},
		{ID: 5, Title: "سبز", RGB: "#00ff00"},
		{ID: 6, Title: "صورتی", RGB: "#ffa500"},
		{ID: 7, Title: "نارنجی", RGB: "#ff4500"},
		{ID: 8, Title: "آبی سماوی", RGB: "#0000ff"},
		{ID: 9, Title: "آبی سماوی", RGB: "#0000ff"},
	}
	err = DB.CreateInBatches(colors, 100).Error
	if err != nil {
		return err
	}
	return nil
}
