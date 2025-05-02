package geoManager

import "github.com/harpy-wings/sol-hr/models"

type GeoManager interface {
	ListStates() ([]*models.MGeoState, error)
	GetState(id int64) (*models.MGeoState, error)

	ListLocations(limit int, offset int, orderBy string) ([]*models.MLocation, error)
	ListLocationsByState(stateID int64, limit int, offset int, orderBy string) ([]*models.MLocation, error)
	QueryLocations(query string, limit int, offset int, orderBy string) ([]*models.MLocation, error)
	CreateLocation(location *models.MLocation) error
	UpdateLocation(location *models.MLocation) error
	DeleteLocation(id int64) error

	GetLocation(id int64) (*models.MLocation, error)
}

var Default GeoManager

func New(opts ...Option) (GeoManager, error) {
	gm := &geoManager{}
	gm.setDefaults()
	for _, opt := range opts {
		err := opt(gm)
		if err != nil {
			return nil, err
		}
	}
	err := gm.init()
	if err != nil {
		return nil, err
	}
	return gm, nil
}

func Init(opts ...Option) error {
	gm, err := New(opts...)
	if err != nil {
		return err
	}
	Default = gm
	return nil
}
