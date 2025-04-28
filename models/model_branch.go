package models

import "errors"

type MBranch struct {
	ID         int64      `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	Title      string     `json:"title" gorm:"type:varchar(255);not null" example:"قرارگاه"`
	Parrent_ID int64      `json:"parrent_id" gorm:"not null" example:"1"`
	Parrent    *MBranch   `json:"parrent" gorm:"foreignKey:Parrent_ID"`
	Children   []*MBranch `json:"children" gorm:"foreignKey:Parrent_ID"`
	Alias      string     `json:"alias" gorm:"type:varchar(255);not null" example:"qrgg"`
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
		{ID: 1, Title: "قرارگاه", Parrent_ID: 0, Alias: ""},
		{ID: 2, Title: "سایت 2", Parrent_ID: 1, Alias: ""},
		{ID: 3, Title: "سایت 3", Parrent_ID: 1, Alias: ""},
		{ID: 4, Title: "کشت و صنعت", Parrent_ID: 0, Alias: ""},
		{ID: 5, Title: "نیرو انسانی", Parrent_ID: 0, Alias: ""},
	}
	err = DB.CreateInBatches(branches, 100).Error
	if err != nil {
		return err
	}
	return nil
}
