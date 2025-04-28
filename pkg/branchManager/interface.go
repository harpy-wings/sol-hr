package branchManager

import "github.com/harpy-wings/sol-hr/models"

type IBranchManager interface {
	ListBranches() ([]*models.MBranch, error)
	ListBranchesByParent(parentID int64) ([]*models.MBranch, error)
	QueryBranches(query string) ([]*models.MBranch, error)
	GetBranch(id int64) (*models.MBranch, error)

	CreateBranch(branch *models.MBranch) error
	UpdateBranch(branch *models.MBranch) error
	DeleteBranch(id int64) error
}
