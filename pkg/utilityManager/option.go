package utilitymanager

import "gorm.io/gorm"

type Option func(*utilityManager) error

func WithDB(db *gorm.DB) Option {
	return func(m *utilityManager) error {
		m.db = db
		return nil
	}
}
