package models

import "errors"

// MAcl represents an Access Control List definition
// @Description Access Control List definition that defines available permissions in the system
type MAcl struct {
	// Unique identifier for the ACL
	ID int64 `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	// Human-readable title of the permission
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"User Management"`
	// Unique key for the permission
	Key string `json:"key" gorm:"type:varchar(255);not null" example:"user-manager"`
}

func (*MAcl) TableName() string {
	return "acl"
}

func (m *MAcl) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MAcl{})
	if err != nil {
		return err
	}
	return nil
}
