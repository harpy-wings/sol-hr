package models

import (
	"errors"

	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	type Seedable interface {
		Seed() error
	}
	seeds := []Seedable{
		&MAcl{},
		&MGeoState{},
		&MAmar{},
		&MBranch{},
		&MLocation{},
		&MMilitaryRank{},
		&MUser{},
		&MSolder{},
		&MSolderBranch{},
		&MExSolder{},
		&MReligon{},
		&MEducationLevel{},
		&MColor{},
		&MUserSession{},
	}

	for _, seed := range seeds {
		if err := seed.Seed(); err != nil {
			return err
		}
	}

	return nil
}
