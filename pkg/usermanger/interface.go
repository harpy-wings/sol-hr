package usermanger

import (
	"github.com/harpy-wings/sol-hr/models"
	"github.com/kataras/iris/v12"
)

type IUserManager interface {
	Login(ctx iris.Context, username string, password string) (u *models.MUser, token string, err error)
	Acl(token string, fnSignature string, permissionType models.PermissionType) (u *models.MUser, err error)
	GetAcls() ([]*models.MAcl, error)
	Logout(token string) error

	ListUsers(req ListUsersRequest) ([]*models.MUser, error)
	GetUser(id int64) (*models.MUser, error)
	GetUserAcls(id int64) ([]*models.MAcl, error)
	UpdateUserAcls(id int64, acls []*models.UserAcl) error
	CreateUser(user *models.MUser) error
	UpdateUser(user *models.MUser) error
	UpdateUserPassword(id int64, password string) error
	DeleteUser(id int64) error
}

type ListUsersRequest struct {
	Limit    int64  `json:"limit" url:"limit"`
	Offset   int64  `json:"offset" url:"offset"`
	Query    string `json:"query" url:"q"`
	BranchID int64  `json:"branch_id" url:"branch_id"`
	OrderBy  string `json:"order_by" url:"order_by"`
}

var Default IUserManager

func New(opts ...Option) (IUserManager, error) {
	um := &userManager{}
	um.setDefaults()
	for _, opt := range opts {
		if err := opt(um); err != nil {
			return nil, err
		}
	}
	err := um.init()
	if err != nil {
		return nil, err
	}
	return um, nil
}

func Init(opts ...Option) error {
	var err error
	Default, err = New(opts...)
	return err
}
