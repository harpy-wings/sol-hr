package models

import (
	"errors"
	"fmt"
	"time"
)

// MUser represents a user in the system
// @Description User model containing all user information including military rank and ACL permissions
type MUser struct {
	// Unique identifier for the user
	ID int64 `json:"id" gorm:"primary_key;auto_increment"`
	// User's first name
	FirstName string `json:"first_name" gorm:"type:varchar(255);not null" example:"John"`
	// User's last name
	LastName string `json:"last_name" gorm:"type:varchar(255);not null" example:"Doe"`
	// Military rank information
	MMilitaryRank *MMilitaryRank `json:"military_rank" gorm:"-"`
	// ID of the user's military rank
	MilitaryRankID int64 `json:"military_rank_id" gorm:"not null" example:"1"`
	// Username for login
	Username string `json:"username" gorm:"type:varchar(255);not null" example:"john.doe"`
	// Password (hashed)
	Password string `json:"password" gorm:"type:varchar(255);not null" example:"secret123"`
	// User role (e.g., admin, user)
	Role string `json:"role" gorm:"type:varchar(255);not null" example:"admin"`

	// System Role
	SystemRole SystemRole `json:"system_role" gorm:"type:int;not null" example:"1"`

	// Super admin flag
	IsSA bool `json:"is_sa" gorm:"type:boolean;default:false" example:"false"`
	// Map of user's ACL permissions
	ACL map[string]*UserAcl `json:"acl" gorm:"-"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" gorm:"not null" example:"2021-01-01T00:00:00Z"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" gorm:"not null" example:"2021-01-01T00:00:00Z"`
	// UID of user who created this user
	CreatedById int64 `json:"created_by_id" gorm:"type:int;not null" example:"1"`
	// Branch IDs
	BranchIDs []int64 `json:"branch_ids" gorm:"-"`
}

func (*MUser) TableName() string {
	return "user"
}

// UserAcl represents Access Control List entries for a user
// @Description Access Control List entry defining a user's permissions for a specific feature
type UserAcl struct {
	// User UID this ACL belongs to
	UserID int64 `json:"user_id" example:"1"`
	// ID of the ACL entry
	AclID int64 `json:"acl_id" example:"1"`
	// Key identifying the permission (e.g., user-manager, acl-management)
	AclKey string `json:"acl_key" example:"user-manager"`
	// Permission value (bitmask: 1=read, 2=write, 4=delete)
	Value int64 `json:"value" example:"7"`
}

func (m *UserAcl) TableName() string {
	return "user_acl"
}

func (m *MUser) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MUser{}, &UserAcl{})
	if err != nil {
		return err
	}
	var users []*MUser
	err = DB.Limit(100).Find(&users).Error
	if err != nil {
		return err
	}
	if len(users) > 0 {
		return nil
	}
	// no users found, create default super admin user
	fmt.Printf("Initialization First SuperAdmin User\nWrite the username:")
	var username string
	fmt.Scanln(&username)
	fmt.Printf("Write the password:")
	var password string
	fmt.Scanln(&password)

	err = DB.Create(&MUser{
		ID:          0,
		Username:    username,
		FirstName:   "Super",
		LastName:    "Admin",
		Password:    password,
		Role:        "superadmin",
		IsSA:        true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedById: 0,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

type SystemRole int64

const (
	SystemRoleUnknown SystemRole = iota
	SystemRoleSuperAdmin
	SystemRoleAdmin
	SystemRoleBranchCommander
	SystemRolePsychologist
)

func (s SystemRole) String() string {
	switch s {
	case SystemRoleUnknown:
		return "unknown"
	case SystemRoleSuperAdmin:
		return "superadmin"
	case SystemRoleAdmin:
		return "admin"
	case SystemRoleBranchCommander:
		return "branch_commander"
	case SystemRolePsychologist:
		return "psychologist"
	default:
		return "unknown"
	}
}
