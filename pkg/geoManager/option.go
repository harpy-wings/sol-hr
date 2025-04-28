package geoManager

import "gorm.io/gorm"

type Option func(*geoManager) error

func WithDB(db *gorm.DB) Option {
	return func(gm *geoManager) error {
		gm.db = db
		return nil
	}
}
