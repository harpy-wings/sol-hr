package branchManager

import "github.com/harpy-wings/sol-hr/models"

type IBranchManager interface {
	ListBranches() ([]*models.MBranch, error)
	ListBranchesByParent(parentID int64) ([]*models.MBranch, error)
	QueryBranches(request QueryBranchesRequest) ([]*models.MBranch, error)
	GetBranch(id int64) (*models.MBranch, error)

	CreateBranch(branch *models.MBranch) error
	UpdateBranch(branch *models.MBranch) error
	DeleteBranch(id int64) error

	ListMilateryBases() ([]*models.MBranch, error)
	GetBranchesByLocation(locationID int64) ([]*models.MBranch, error)
}

var Default IBranchManager

func New(opts ...Option) (IBranchManager, error) {
	m := &branchManager{}
	m.setDefaults()
	for _, opt := range opts {
		opt(m)
	}
	err := m.init()
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Init(opts ...Option) error {
	m, err := New(opts...)
	if err != nil {
		return err
	}
	Default = m
	return nil
}
