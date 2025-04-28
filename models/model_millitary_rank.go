package models

import "errors"

// MMilitaryRank represents a military rank in the system
// @Description Military rank definition with title and logo
type MMilitaryRank struct {
	// Unique identifier for the military rank
	ID int64 `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	// Title/Name of the military rank
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"سرهنگ"`
	// URL to the rank's logo/insignia
	Logo string `json:"logo" gorm:"type:varchar(255);not null" example:"https://via.placeholder.com/150"`
}

func (*MMilitaryRank) TableName() string {
	return "military_rank"
}

func (m *MMilitaryRank) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MMilitaryRank{})
	if err != nil {
		return err
	}
	eranks := []MMilitaryRank{}
	err = DB.Find(&eranks).Error
	if err != nil {
		return err
	}
	if len(eranks) > 0 {
		return nil
	}
	var ranks = []MMilitaryRank{
		{
			Title: "سرباز",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرباز دوم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرباز یکم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرجوخه",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "گروهبان سوم",
			Logo:  "https://via.placeholder.com/150",
		}, {
			Title: "گروهبان دوم",
			Logo:  "https://via.placeholder.com/150",
		}, {
			Title: "گروهبان یکم",
			Logo:  "https://via.placeholder.com/150",
		}, {
			Title: "استوار دوم",
			Logo:  "https://via.placeholder.com/150",
		}, {
			Title: "استوار یکم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "ستوان سوم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "ستوان دوم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "ستوان یکم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سروان",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرگرد",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرهنگ دوم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرهنگ",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرتیپ دوم",
			Logo:  "https://via.placeholder.com/150",
		},
		{
			Title: "سرتیپ",
			Logo:  "https://via.placeholder.com/150",
		},
	}
	err = DB.CreateInBatches(ranks, 100).Error
	if err != nil {
		return err
	}
	return nil
}
