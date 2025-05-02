package solderManager

import (
	"github.com/harpy-wings/sol-hr/pkg/branchManager"
	"github.com/harpy-wings/sol-hr/pkg/geoManager"
	utilitymanager "github.com/harpy-wings/sol-hr/pkg/utilityManager"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Option func(*solderManager) error

func WithLogger(logger logrus.FieldLogger) Option {
	return func(m *solderManager) error {
		m.logger = logger
		return nil
	}
}

func WithDB(db *gorm.DB) Option {
	return func(m *solderManager) error {
		m.db = db
		return nil
	}
}
func WithGeoManager(geoManager geoManager.GeoManager) Option {
	return func(m *solderManager) error {
		m.geoManager = geoManager
		return nil
	}
}

func WithBranchManager(branchManager branchManager.IBranchManager) Option {
	return func(m *solderManager) error {
		m.branchManager = branchManager
		return nil
	}
}

func WithUtilityManager(utilityManager utilitymanager.IUtilityManager) Option {
	return func(m *solderManager) error {
		m.utilityManager = utilityManager
		return nil
	}
}
