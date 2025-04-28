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
