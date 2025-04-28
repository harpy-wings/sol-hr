package user_controller

import (
	"errors"
	"strings"
	"time"

	"github.com/harpy-wings/sol-hr/controllers"
	"github.com/harpy-wings/sol-hr/models"
	"github.com/harpy-wings/sol-hr/pkg/usermanger"
	"github.com/kataras/iris/v12"
)

func (c *controller) Get(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("token is required")
	}
	token = strings.TrimPrefix(token, "Bearer ")
	user, err := c.userService.Acl(token, "", 0)
	if err != nil {
		return nil, err
	}
	user, err = c.userService.GetUser(user.ID)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Data: user}, nil
}

func (c *controller) GetList(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("token is required")
	}
	var req usermanger.ListUsersRequest
	err := ctx.ReadQuery(&req)
	if err != nil {
		return nil, err
	}
	token = strings.TrimPrefix(token, "Bearer ")
	_, err = c.userService.Acl(token, "user-management", models.PermissionTypeReadOnly)
	if err != nil {
		return nil, err
	}
	users, err := c.userService.ListUsers(req)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Data: users}, nil
}

func (c *controller) GetAcls(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("token is required")
	}

	token = strings.TrimPrefix(token, "Bearer ")
	_, err := c.userService.Acl(token, "user-management", models.PermissionTypeReadOnly)
	if err != nil {
		return nil, err
	}
	acls, err := c.userService.GetAcls()
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Data: acls}, nil
}

func (c *controller) Post(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("token is required")
	}

	token = strings.TrimPrefix(token, "Bearer ")
	currentUser, err := c.userService.Acl(token, "user-management", models.PermissionTypeReadWrite)
	if err != nil {
		return nil, err
	}

	var user models.MUser
	err = ctx.ReadJSON(&user)
	if err != nil {
		return nil, err
	}
	user.CreatedBy = currentUser.ID
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = c.userService.CreateUser(&user)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Data: user}, nil
}

func (c *controller) Put(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("token is required")
	}
	var user models.MUser
	err := ctx.ReadJSON(&user)
	if err != nil {
		return nil, err
	}
	token = strings.TrimPrefix(token, "Bearer ")
	currentUser, err := c.userService.Acl(token, "", 0)
	if err != nil {
		return nil, err
	}

	if user.ID != currentUser.ID {
		_, err := c.userService.Acl(token, "user-management", models.PermissionTypeReadWrite)
		if err != nil {
			return nil, err
		}
	}

	err = c.userService.UpdateUser(&user)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Data: user}, nil
}

func (c *controller) PutByPassword(ctx iris.Context, id int64) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("token is required")
	}
	_, err := c.userService.Acl(token, "user-management", models.PermissionTypeReadWrite)
	if err != nil {
		return nil, err
	}
	var req struct {
		Password string `json:"password"`
	}
	err = ctx.ReadJSON(&req)
	if err != nil {
		return nil, err
	}
	err = c.userService.UpdateUserPassword(id, req.Password)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Data: "Password updated successfully"}, nil
}

func (c *controller) DeleteBy(ctx iris.Context, id int64) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		return nil, errors.New("token is required")
	}
	_, err := c.userService.Acl(token, "user-management", models.PermissionTypeDelete)
	if err != nil {
		return nil, err
	}
	err = c.userService.DeleteUser(id)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Data: "User deleted successfully"}, nil
}
