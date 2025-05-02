package utilitymanager

import (
	"errors"

	"github.com/harpy-wings/sol-hr/models"
	"gorm.io/gorm"
)

type utilityManager struct {
	db *gorm.DB

	// Cache
	militaryRanks     map[int64]*models.MMilitaryRank
	militaryRanksList []*models.MMilitaryRank
}

func (m *utilityManager) setDefaults() {
	if models.DB != nil {
		m.db = models.DB
	}
}

func (m *utilityManager) init() error {
	if m.db == nil {
		return errors.New("database is not initialized")
	}

	if m.militaryRanks == nil {
		m.militaryRanks = make(map[int64]*models.MMilitaryRank)
	}
	err := m.loadMilitaryRanks()
	if err != nil {
		return err
	}
	return nil
}

func (m *utilityManager) ListMilitaryRanks() ([]*models.MMilitaryRank, error) {
	return m.militaryRanksList, nil
}

func (m *utilityManager) GetMilitaryRank(id int64) (*models.MMilitaryRank, error) {
	if m.militaryRanks == nil {
		m.militaryRanks = make(map[int64]*models.MMilitaryRank)
	}
	rank, ok := m.militaryRanks[id]
	if !ok {
		return nil, errors.New("military rank not found")
	}
	return rank, nil
}
func (m *utilityManager) loadMilitaryRanks() error {
	m.militaryRanksList = make([]*models.MMilitaryRank, 0)
	m.militaryRanks = make(map[int64]*models.MMilitaryRank)

	var ranks []*models.MMilitaryRank
	err := m.db.Find(&ranks).Error
	if err != nil {
		return err
	}

	for _, rank := range ranks {
		m.militaryRanksList = append(m.militaryRanksList, rank)
		m.militaryRanks[rank.ID] = rank
	}

	return nil
}
