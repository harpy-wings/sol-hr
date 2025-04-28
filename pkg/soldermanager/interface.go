package solderManager

import "github.com/harpy-wings/sol-hr/models"

type ISolderManager interface {
	ListSolder(req ListSolderRequest) ([]*models.MSolder, error)

	GetSolder(id int64) (*models.MSolder, error)
	CreateSolder(solder *models.MSolder) error
	UpdateSolder(solder *models.MSolder) error
	DeleteSolder(id int64) error
}

type ListSolderRequest struct {
	Limit    int64  `json:"limit" url:"limit"`
	Offset   int64  `json:"offset" url:"offset"`
	Query    string `json:"query" url:"q"`
	BranchID int64  `json:"branch_id" url:"branch_id"`
}
