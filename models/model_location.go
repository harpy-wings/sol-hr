package models

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

// MLocation represents a location in the system
// @Description Location model containing geographical and environmental information
type MLocation struct {
	// Unique identifier for the location
	ID int64 `json:"id" gorm:"primary_key;auto_increment" example:"1"`
	// Name/Title of the location
	Title string `json:"title" gorm:"type:varchar(255);not null" example:"تهران"`
	// Associated state information
	State MGeoState `json:"state" gorm:"-"`
	// Reference to the state this location belongs to
	StateID int64 `json:"state_id" gorm:"not null" example:"1"`
	// Latitude coordinate
	Lat float64 `json:"lat" gorm:"type:float;not null" example:"35.6892"`
	// Longitude coordinate
	Lng float64 `json:"lng" gorm:"type:float;not null" example:"51.3890"`

	// Distance from the reference location (used for calculations)
	DistanceKm float64 `json:"distance_km" gorm:"type:float;" example:"100"`
}

func (*MLocation) TableName() string {
	return "location"
}

func (m *MLocation) Seed() error {
	if DB == nil {
		return errors.New("database is not initialized")
	}
	err := DB.AutoMigrate(&MLocation{})
	if err != nil {
		return err
	}
	locations := []*MLocation{}
	err = DB.Find(&locations).Error
	if err != nil {
		return err
	}
	if len(locations) > 0 {
		return nil
	}
	csvFile, err := os.Open("./data/locations.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	records, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	var MLocations []*MLocation

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
		StateID, err := strconv.ParseInt(record[2], 10, 64)
		if err != nil {
			return err
		}
		Lat, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return err
		}
		Lng, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			return err
		}

		DistanceKm, err := strconv.ParseFloat(record[7], 64)
		if err != nil {
			return err
		}
		MLocations = append(MLocations, &MLocation{
			ID:         ID,
			Title:      Title,
			StateID:    StateID,
			Lat:        Lat,
			Lng:        Lng,
			DistanceKm: DistanceKm,
		})
	}

	err = DB.CreateInBatches(MLocations, 100).Error
	if err != nil {
		return err
	}
	return nil

}
