package geoManager

import (
	"github.com/harpy-wings/sol-hr/models"
	"gorm.io/gorm"
)

type geoManager struct {
	db *gorm.DB

	// caches
	states        map[int64]*models.MGeoState
	statesList    []*models.MGeoState
	locations     map[int64]*models.MLocation
	locationsList []*models.MLocation
}

var _ GeoManager = &geoManager{}

func (gm *geoManager) setDefaults() {
	if models.DB != nil {
		gm.db = models.DB
	}
}

func (gm *geoManager) init() error {
	err := gm.loadStates()
	if err != nil {
		return err
	}
	err = gm.loadLocations()
	if err != nil {
		return err
	}
	return nil
}

func (m *geoManager) ListStates() ([]*models.MGeoState, error) {
	return m.statesList, nil
}
func (m *geoManager) GetState(id int64) (*models.MGeoState, error) {
	return m.states[id], nil
}

func (m *geoManager) ListLocations(limit int, offset int, orderBy string) ([]*models.MLocation, error) {
	var locations []*models.MLocation
	if limit == 0 {
		limit = 100
	}
	if offset == 0 {
		offset = 0
	}
	if orderBy == "" {
		orderBy = "created_at DESC"
	}
	err := m.db.Limit(int(limit)).Offset(int(offset)).Order(orderBy).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	for _, location := range locations {
		state, ok := m.states[location.StateID]
		if !ok {
			continue
		}
		location.State = *state
	}
	return locations, nil
}

func (m *geoManager) ListLocationsByState(stateID int64, limit int, offset int, orderBy string) ([]*models.MLocation, error) {
	var locations []*models.MLocation
	if limit == 0 {
		limit = 100
	}
	if offset == 0 {
		offset = 0
	}
	if orderBy == "" {
		orderBy = "created_at DESC"
	}
	err := m.db.Limit(int(limit)).Offset(int(offset)).Order(orderBy).Where("state_id = ?", stateID).Find(&locations).Error
	if err != nil {
		return nil, err
	}
	for _, location := range locations {
		state, ok := m.states[location.StateID]
		if !ok {
			continue
		}
		location.State = *state
	}
	return locations, nil
}

func (m *geoManager) QueryLocations(query string, limit int, offset int, orderBy string) ([]*models.MLocation, error) {
	var locations []*models.MLocation
	if limit == 0 {
		limit = 100
	}
	if offset == 0 {
		offset = 0
	}
	if orderBy == "" {
		orderBy = "created_at DESC"
	}
	err := m.db.Limit(int(limit)).Offset(int(offset)).Order(orderBy).Where("title LIKE ?", "%"+query+"%").Find(&locations).Error
	if err != nil {
		return nil, err
	}
	for _, location := range locations {
		state, ok := m.states[location.StateID]
		if !ok {
			continue
		}
		location.State = *state
	}
	return locations, nil
}

func (m *geoManager) GetLocation(id int64) (*models.MLocation, error) {
	return m.locations[id], nil
}

func (m *geoManager) CreateLocation(location *models.MLocation) error {
	return m.db.Create(location).Error
}

func (m *geoManager) UpdateLocation(location *models.MLocation) error {
	return m.db.Save(location).Error
}

func (m *geoManager) DeleteLocation(id int64) error {
	err := m.db.Delete(&models.MLocation{}, id).Error
	if err != nil {
		return err
	}
	delete(m.locations, id)
	return nil
}

//#region Helpers

func (m *geoManager) loadStates() error {
	var states []*models.MGeoState
	err := m.db.Find(&states).Error
	if err != nil {
		return err
	}
	m.states = make(map[int64]*models.MGeoState)
	for _, state := range states {
		m.states[state.ID] = state
	}
	m.statesList = states
	return nil
}

func (m *geoManager) loadLocations() error {
	var locations []*models.MLocation
	err := m.db.Find(&locations).Error
	if err != nil {
		return err
	}
	m.locations = make(map[int64]*models.MLocation)
	for _, location := range locations {
		m.locations[location.ID] = location
	}
	m.locationsList = locations
	return nil
}

//#endregion
