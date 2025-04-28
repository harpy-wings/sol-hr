package utilitymanager

import "github.com/harpy-wings/sol-hr/models"

type IUtilityManager interface {
	ListMilitaryRanks() ([]*models.MMilitaryRank, error)
	GetMilitaryRank(id int64) (*models.MMilitaryRank, error)
}

var Default IUtilityManager

func New(opts ...Option) (IUtilityManager, error) {
	um := &utilityManager{}
	um.setDefaults()
	for _, opt := range opts {
		err := opt(um)
		if err != nil {
			return nil, err
		}
	}
	err := um.init()
	if err != nil {
		return nil, err
	}
	return um, nil
}
