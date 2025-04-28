package geoManager

import (
	"errors"

	"github.com/harpy-wings/sol-hr/models"
	"gorm.io/gorm"
)

type geoManager struct {
	db *gorm.DB

	// caches
	states            map[int64]*models.MGeoState
	statesList        []*models.MGeoState
	locations         map[int64]*models.MLocation
	locationsList     []*models.MLocation
	milateryBases     map[int64]*models.MMilateryBase
	milateryBasesList []*models.MMilateryBase
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
	err = gm.loadMilateryBases()
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
	err := m.db.Limit(int(limit)).Offset(int(offset)).Order(orderBy).Where("name LIKE ?", "%"+query+"%").Find(&locations).Error
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

func (m *geoManager) ListMilateryBases() ([]*models.MMilateryBase, error) {
	return m.milateryBasesList, nil
}

func (m *geoManager) QueryMilateryBases(query string, limit int, offset int, orderBy string) ([]*models.MMilateryBase, error) {
	var milateryBases []*models.MMilateryBase
	err := m.db.Transaction(func(tx *gorm.DB) error {
		if limit == 0 {
			limit = 100
		}
		if offset == 0 {
			offset = 0
		}
		if orderBy == "" {
			orderBy = "created_at DESC"
		}
		err := tx.Limit(int(limit)).Offset(int(offset)).Order(orderBy).Where("name LIKE ?", "%"+query+"%").Find(&milateryBases).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	for _, milateryBase := range milateryBases {
		location, ok := m.locations[milateryBase.LocationID]
		if !ok {
			continue
		}
		milateryBase.Location = location
	}
	return milateryBases, nil
}

func (m *geoManager) GetMilateryBase(id int64) (*models.MMilateryBase, error) {
	return m.milateryBases[id], nil
}
func (m *geoManager) GetMilateryBaseByLocation(locationID int64) (*models.MMilateryBase, error) {
	for _, milateryBase := range m.milateryBasesList {
		if milateryBase.LocationID == locationID {
			return milateryBase, nil
		}
	}
	return nil, errors.New("milatery base not found")
}
func (m *geoManager) CreateMilateryBase(milateryBase *models.MMilateryBase) error {
	err := m.db.Create(milateryBase).Error
	if err != nil {
		return err
	}
	m.milateryBases[milateryBase.ID] = milateryBase
	return nil
}
func (m *geoManager) UpdateMilateryBase(milateryBase *models.MMilateryBase) error {
	err := m.db.Save(milateryBase).Error
	if err != nil {
		return err
	}
	m.milateryBases[milateryBase.ID] = milateryBase
	return nil
}
func (m *geoManager) DeleteMilateryBase(id int64) error {
	err := m.db.Delete(&models.MMilateryBase{}, id).Error
	if err != nil {
		return err
	}
	delete(m.milateryBases, id)
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

func (m *geoManager) loadMilateryBases() error {
	var milateryBases []*models.MMilateryBase
	err := m.db.Find(&milateryBases).Error
	if err != nil {
		return err
	}
	m.milateryBases = make(map[int64]*models.MMilateryBase)
	for _, milateryBase := range milateryBases {
		m.milateryBases[milateryBase.ID] = milateryBase
	}
	m.milateryBasesList = milateryBases
	return nil
}

//#endregion
