package adjustmentManager

import "github.com/harpy-wings/sol-hr/models"

type IAdjustmentManager interface {
	ListAdjustments(solderUid string) ([]*models.MAdjustment, error)
	GetAdjustment(id int64) (*models.MAdjustment, error)

	CreateAdjustment(adjustment *models.MAdjustment) error
	UpdateAdjustment(adjustment *models.MAdjustment) error
	DeleteAdjustment(id int64) error
}

var Default IAdjustmentManager
