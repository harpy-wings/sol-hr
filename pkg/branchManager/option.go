package branchManager

import (
	"github.com/harpy-wings/sol-hr/pkg/geoManager"
	"gorm.io/gorm"
)

type Option func(*branchManager) error

func WithDB(db *gorm.DB) Option {
	return func(m *branchManager) error {
		m.db = db
		return nil
	}
}
func WithGeoManager(geoManager geoManager.GeoManager) Option {
	return func(m *branchManager) error {
		m.geoManager = geoManager
		return nil
	}
}
