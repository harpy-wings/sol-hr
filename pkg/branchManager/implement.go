package branchManager

import (
	"errors"

	"github.com/harpy-wings/sol-hr/models"
	"github.com/harpy-wings/sol-hr/pkg/geoManager"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type branchManager struct {
	logger        logrus.FieldLogger
	db            *gorm.DB
	geoManager    geoManager.GeoManager
	branchesList  []*models.MBranch
	rootBranches  []*models.MBranch
	milateryBases []*models.MBranch
	branches      map[int64]*models.MBranch
}

func (m *branchManager) setDefaults() {
	m.logger = logrus.StandardLogger()
	if models.DB != nil {
		m.db = models.DB
	}
}

func (m *branchManager) init() error {
	if m.db == nil {
		return errors.New("db is not set")
	}
	err := m.loadBranches()
	if err != nil {
		return err
	}
	return nil
}

func (m *branchManager) ListBranches() ([]*models.MBranch, error) {
	if m.rootBranches == nil {
		err := m.loadBranches()
		if err != nil {
			m.logger.Error(err)
			return nil, err
		}
	}
	return m.rootBranches, nil
}

func (m *branchManager) ListBranchesByParent(parentID int64) ([]*models.MBranch, error) {
	var branches []*models.MBranch
	err := m.db.Where("parrent_id = ?", parentID).Find(&branches).Error
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}
	m.enrichBranchParent(branches)
	m.enrichBranchChildren(branches)
	return branches, nil
}

func (m *branchManager) QueryBranches(request QueryBranchesRequest) ([]*models.MBranch, error) {
	var branches []*models.MBranch
	err := m.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Model(&models.MBranch{})
		if request.Xid != nil {
			tx = tx.Where("xid = ?", request.Xid)
		}
		if request.Query != "" {
			tx = tx.Where("title LIKE ? OR alias LIKE ?", "%"+request.Query+"%", "%"+request.Query+"%")
		}
		if request.OrderBy != "" {
			tx = tx.Order(request.OrderBy)
		} else {
			tx = tx.Order("created_at DESC")
		}
		err := tx.Find(&branches).Error
		if err != nil {
			m.logger.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		m.logger.Error(err)
		return nil, err
	}
	m.enrichBranchParent(branches)
	m.enrichBranchChildren(branches)
	return branches, nil
}

func (m *branchManager) GetBranch(id int64) (*models.MBranch, error) {
	branch, ok := m.branches[id]
	if !ok {
		m.logger.Error("branch not found")
		return nil, errors.New("branch not found")
	}
	return branch, nil
}

func (m *branchManager) CreateBranch(branch *models.MBranch) error {
	if branch.LocationID != 0 {
		location, err := m.geoManager.GetLocation(branch.LocationID)
		if err != nil {
			m.logger.Error(err)
			return err
		}
		branch.Location = location
	}
	err := m.db.Create(branch).Error
	if err != nil {
		m.logger.Error(err)
		return err
	}
	m.branches[branch.ID] = branch
	m.loadBranches()
	return nil
}

func (m *branchManager) UpdateBranch(branch *models.MBranch) error {
	if branch.ID == 0 {
		return errors.New("شناسه برای بروز رسانی یگان یا قسمت مورد نظر است")
	}
	err := m.db.Save(branch).Error
	if err != nil {
		return err
	}
	m.branches[branch.ID] = branch
	m.loadBranches()
	return nil
}

func (m *branchManager) DeleteBranch(id int64) error {
	err := m.db.Delete(&models.MBranch{}, id).Error
	if err != nil {
		return err
	}
	m.loadBranches()
	return nil
}

func (m *branchManager) ListMilateryBases() ([]*models.MBranch, error) {
	return m.milateryBases, nil
}

func (m *branchManager) GetBranchesByLocation(locationID int64) ([]*models.MBranch, error) {
	var branches []*models.MBranch
	for _, branch := range m.branchesList {
		if branch.LocationID == locationID {
			branches = append(branches, branch)
		}
	}
	return branches, nil
}

// Internal functions
func (m *branchManager) loadBranches() error {
	var branches []*models.MBranch
	var rootBranches []*models.MBranch
	err := m.db.Find(&branches).Error
	if err != nil {
		m.logger.Error(err)
		return err
	}

	m.branches = make(map[int64]*models.MBranch)
	for _, branch := range branches {
		m.branches[branch.ID] = branch
	}

	for i, branch := range branches {
		if branch.Xid == 0 {
			if branch.ParrentID != 0 {
				m.branches[branch.ParrentID].Children = append(m.branches[branch.ParrentID].Children, branch)
			}
			branches[i].Parrent = m.branches[branch.ParrentID]
			if branch.LocationID != 0 {
				location, err := m.geoManager.GetLocation(branch.LocationID)
				if err != nil {
					m.logger.Error(err)
					return err
				}
				branches[i].Location = location
			}
			if branch.ParrentID == 0 {
				rootBranches = append(rootBranches, branch)
			}
		} else {
			m.milateryBases = append(m.milateryBases, branch)
		}

	}

	m.branchesList = branches
	m.rootBranches = rootBranches
	return nil
}

func (m *branchManager) enrichBranchParent(branch []*models.MBranch) {
	for i, b := range branch {
		if b.ParrentID != 0 {
			branch[i].Parrent = m.branches[b.ParrentID]
		}
	}
}

func (m *branchManager) enrichBranchChildren(branches []*models.MBranch) {
	for i, b := range branches {
		branches[i].Children = m.branches[b.ID].Children
	}
}
