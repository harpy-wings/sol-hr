package models

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

// MGeoState represents a geographical state/province
// @Description Geographical state or province definition used for location categorization
type MGeoState struct {
	// Unique identifier for the state
	ID int64 `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	// Name of the state/province
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"تهران"`
}

func (*MGeoState) TableName() string {
	return "geo_state"
}

func (m *MGeoState) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MGeoState{})
	if err != nil {
		return err
	}
	estates := []MGeoState{}
	err = DB.Find(&estates).Error
	if err != nil {
		return err
	}
	if len(estates) > 0 {
		return nil
	}
	csvFile, err := os.Open("./data/geo_states.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	for i, record := range records {
		if i == 0 {
			// skip header
			continue
		}
		ID, err := strconv.ParseInt(record[0], 10, 64)
		if err != nil {
			return err
		}
		Title := record[1]
		DB.Create(&MGeoState{ID: ID, Title: Title})
	}
	return nil
}
