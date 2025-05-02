package solderManager

import (
	"time"

	"github.com/harpy-wings/sol-hr/models"
)

type ISolderManager interface {
	ListSolder(req ListSolderRequest) ([]*models.MSolder, error)

	GetSolder(uid string) (*models.MSolder, error)
	CreateSolder(solder *models.MSolder) error
	UpdateSolder(solder *models.MSolder) error
	DeleteSolder(uid string) error

	// AlocateBranch is used to alocate a branch to a solder
	AlocateBranch(req AlocateBranchRequest) error
}

type ListSolderRequest struct {
	Limit  int    `json:"limit" url:"limit"`
	Offset int    `json:"offset" url:"offset"`
	Query  string `json:"query" url:"query"`

	FirstName  string `json:"first_name" url:"first_name"`
	LastName   string `json:"last_name" url:"last_name"`
	FatherName string `json:"father_name" url:"father_name"`

	BranchID []int64 `json:"branch_id" url:"branch_id"`
	OrderBy  string  `json:"order_by" url:"order_by"`
	Status   *int64  `json:"status" url:"status"`

	ServiceStartedAt *time.Time `json:"service_started_at" url:"service_started_at"`
	ServiceEndAt     *time.Time `json:"service_end_at" url:"service_end_at"`

	HasDisability     *bool `json:"has_disability" url:"has_disability"`
	IsMentallyHealthy *bool `json:"is_mentally_healthy" url:"is_mentally_healthy"`
}

var Default ISolderManager

func New(opts ...Option) (ISolderManager, error) {
	m := &solderManager{}
	m.setDefaults()
	for _, opt := range opts {
		err := opt(m)
		if err != nil {
			return nil, err
		}
	}
	if err := m.init(); err != nil {
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
