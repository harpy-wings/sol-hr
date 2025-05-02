package solderManager

import (
	"errors"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/harpy-wings/sol-hr/models"
	"github.com/harpy-wings/sol-hr/pkg/branchManager"
	"github.com/harpy-wings/sol-hr/pkg/geoManager"
	utilitymanager "github.com/harpy-wings/sol-hr/pkg/utilityManager"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type solderManager struct {
	logger logrus.FieldLogger
	db     *gorm.DB

	geoManager     geoManager.GeoManager
	branchManager  branchManager.IBranchManager
	utilityManager utilitymanager.IUtilityManager
}

func (m *solderManager) setDefaults() {
	m.logger = logrus.StandardLogger()
	if models.DB != nil {
		m.db = models.DB
	}
}

func (m *solderManager) init() error {
	if m.db == nil {
		return errors.New("پایگاه داده یافت نشد")
	}
	return nil
}

func (m *solderManager) ListSolder(req ListSolderRequest) ([]*models.MSolder, error) {
	var solder []*models.MSolder
	err := m.db.Transaction(func(tx *gorm.DB) error {
		tx = tx.Model(&models.MSolder{})
		if req.Status != nil {
			tx = tx.Where("status = ?", req.Status)
		}
		if req.BranchID != nil {
			tx = tx.Where("( primary_branch_id IN ? OR secondary_branch_id IN ? )", req.BranchID, req.BranchID)
		}
		if req.Query != "" {
			tx = tx.Where("first_name LIKE ? OR last_name LIKE ? OR father_name LIKE ? OR uid LIKE ? OR personel_id LIKE ?", "%"+req.Query+"%", "%"+req.Query+"%", "%"+req.Query+"%", "%"+req.Query+"%", "%"+req.Query+"%")
		}
		if req.FirstName != "" {
			tx = tx.Where("first_name LIKE ?", "%"+req.FirstName+"%")
		}
		if req.LastName != "" {
			tx = tx.Where("last_name LIKE ?", "%"+req.LastName+"%")
		}
		if req.FatherName != "" {
			tx = tx.Where("father_name LIKE ?", "%"+req.FatherName+"%")
		}

		if req.ServiceStartedAt != nil {
			tx = tx.Where("service_started_at >= ?", req.ServiceStartedAt)
		}
		if req.ServiceEndAt != nil {
			tx = tx.Where("discharge_date <= ?", req.ServiceEndAt)
		}
		if req.HasDisability != nil {
			tx = tx.Where("has_disability = ?", req.HasDisability)
		}
		if req.IsMentallyHealthy != nil {
			tx = tx.Where("is_mentally_healthy = ?", req.IsMentallyHealthy)
		}

		if req.OrderBy != "" {
			tx = tx.Order(req.OrderBy)
		} else {
			tx = tx.Order("updated_at DESC")
		}
		if req.Limit != 0 {
			tx = tx.Limit(req.Limit)
		} else {
			tx = tx.Limit(100)
		}

		if req.Offset != 0 {
			tx = tx.Offset(req.Offset)
		} else {
			tx = tx.Offset(0)
		}

		if err := tx.Find(&solder).Error; err != nil {
			return err
		}
		return nil
	})

	for _, solder := range solder {
		if solder.PrimaryBranchID != 0 {
			solder.PrimaryBranch, err = m.branchManager.GetBranch(solder.PrimaryBranchID)
			if err != nil {
				return nil, err
			}
		}
		if solder.SecondaryBranchID != 0 {
			solder.SecondaryBranch, err = m.branchManager.GetBranch(solder.SecondaryBranchID)
			if err != nil {
				return nil, err
			}
		}
		solder.MilitaryRank, err = m.utilityManager.GetMilitaryRank(solder.MilitaryRankID)
		if err != nil {
			return nil, err
		}

	}
	return solder, err
}

func (m *solderManager) GetSolder(uid string) (*models.MSolder, error) {
	var solder models.MSolder
	err := m.db.First(&solder, "uid = ?", uid).Error
	if err != nil {
		return nil, err
	}
	solder.MilitaryRank, err = m.utilityManager.GetMilitaryRank(solder.MilitaryRankID)
	if err != nil {
		return nil, err
	}
	if solder.PrimaryBranchID != 0 {
		solder.PrimaryBranch, err = m.branchManager.GetBranch(solder.PrimaryBranchID)
		if err != nil {
			return nil, err
		}
	}
	if solder.SecondaryBranchID != 0 {
		solder.SecondaryBranch, err = m.branchManager.GetBranch(solder.SecondaryBranchID)
		if err != nil {
			return nil, err
		}
	}
	// todo load leave profile
	// todo load ditails

	return &solder, err
}

func (m *solderManager) CreateSolder(solder *models.MSolder) error {
	validator := validator.New()
	err := validator.Struct(solder)
	if err != nil {
		return err
	}
	solder.CreatedAt = time.Now()
	solder.UpdatedAt = time.Now()
	// if primary branch is set, alocate branch

	err = m.db.Create(solder).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *solderManager) UpdateSolder(solder *models.MSolder) error {
	validator := validator.New()
	err := validator.Struct(solder)
	if err != nil {
		return err
	}
	solder.UpdatedAt = time.Now()
	err = m.db.Save(solder).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *solderManager) DeleteSolder(uid string) error {
	// todo: delete from everywhere and Branch Statistics and Leave Profile and leaves and etc.
	err := m.db.Delete(&models.MSolder{}, "uid = ?", uid).Error
	if err != nil {
		return err
	}
	return nil
}

func (m *solderManager) AlocateBranch(req AlocateBranchRequest) error {
	return errors.New("not implemented")
}
func (m *solderManager) DeleteAlocatedBranch(id int64) error {
	return errors.New("not implemented")
}
