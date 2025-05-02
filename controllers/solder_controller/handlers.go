package solder_controller

import (
	"errors"

	"github.com/harpy-wings/sol-hr/controllers"
	"github.com/harpy-wings/sol-hr/models"
	solderManager "github.com/harpy-wings/sol-hr/pkg/soldermanager"
	"github.com/kataras/iris/v12"
)

func (c *controller) GetList(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	user, err := c.userService.Acl(token, "solder", models.PermissionTypeReadOnly)
	if err != nil {
		return nil, err
	}
	var req solderManager.ListSolderRequest
	err = ctx.ReadQuery(&req)
	if err != nil {
		return nil, err
	}

	if !user.IsSA {
		req.BranchID = user.BranchIDs
	}

	solders, err := c.solderService.ListSolder(req)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: solders}, nil
}

func (c *controller) Post(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	user, err := c.userService.Acl(token, "solder", models.PermissionTypeReadWrite)
	if err != nil {
		return nil, err
	}
	if user.SystemRole != models.SystemRoleAdmin && !user.IsSA {
		return nil, errors.New("شما نمیتوانید این داده را ایجاد کنید")
	}

	var solder models.MSolder
	err = ctx.ReadJSON(&solder)
	if err != nil {
		return nil, err
	}
	solder.CreatedByID = user.ID
	err = c.solderService.CreateSolder(&solder)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: solder}, nil
}

func (c *controller) Put(ctx iris.Context) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	user, err := c.userService.Acl(token, "solder", models.PermissionTypeReadWrite)
	if err != nil {
		return nil, err
	}

	if user.SystemRole != models.SystemRoleAdmin && !user.IsSA {
		return nil, errors.New("شما نمیتوانید این داده را ویرایش کنید")
	}

	var solder models.MSolder
	err = ctx.ReadJSON(&solder)
	if err != nil {
		return nil, err
	}

	err = c.solderService.UpdateSolder(&solder)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: solder}, nil
}

func (c *controller) DeleteBy(ctx iris.Context, uid string) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	user, err := c.userService.Acl(token, "solder", models.PermissionTypeDelete)
	if err != nil {
		return nil, err
	}
	if user.SystemRole != models.SystemRoleAdmin && !user.IsSA {
		return nil, errors.New("شما نمیتوانید این داده را حذف کنید")
	}

	err = c.solderService.DeleteSolder(uid)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true}, nil
}

func (c *controller) GetBy(ctx iris.Context, uid string) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	_, err := c.userService.Acl(token, "solder", models.PermissionTypeReadOnly)
	if err != nil {
		return nil, err
	}

	solder, err := c.solderService.GetSolder(uid)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: solder}, nil
}
func (c *controller) Options(ctx iris.Context) {
	ctx.StatusCode(200)
}

func (c *controller) OptionsBy(ctx iris.Context) {
	ctx.StatusCode(200)
}
