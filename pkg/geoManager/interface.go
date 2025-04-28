package geoManager

import "github.com/harpy-wings/sol-hr/models"

type GeoManager interface {
	ListStates() ([]*models.MGeoState, error)
	GetState(id int64) (*models.MGeoState, error)

	ListLocations(limit int64, offset int64, orderBy string) ([]*models.MLocation, error)
	ListLocationsByState(stateID int64, limit int64, offset int64, orderBy string) ([]*models.MLocation, error)
	QueryLocations(query string, limit int64, offset int64, orderBy string) ([]*models.MLocation, error)
	CreateLocation(location *models.MLocation) error
	UpdateLocation(location *models.MLocation) error
	DeleteLocation(id int64) error

	ListMilateryBases() ([]*models.MMilateryBase, error)
	QueryMilateryBases(query string, limit int64, offset int64, orderBy string) ([]*models.MMilateryBase, error)
	GetMilateryBase(id int64) (*models.MMilateryBase, error)
	GetMilateryBaseByLocation(locationID int64) (*models.MMilateryBase, error)
	CreateMilateryBase(milateryBase *models.MMilateryBase) error
	UpdateMilateryBase(milateryBase *models.MMilateryBase) error
	DeleteMilateryBase(id int64) error
}
