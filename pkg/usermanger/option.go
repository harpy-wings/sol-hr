package usermanger

import "gorm.io/gorm"

type Option func(*userManager) error

func WithDB(db *gorm.DB) Option {
	return func(u *userManager) error {
		u.db = db
		return nil
	}
}
