package branch_controller

import (
	"errors"
	"time"

	"github.com/harpy-wings/sol-hr/controllers"
	"github.com/harpy-wings/sol-hr/models"
	"github.com/kataras/iris/v12"
)

// /api/branches/list
// /api/branches/bases

func (c *controller) GetList(ctx iris.Context) (*controllers.GenericResponse, error) {

	branches, err := c.branchService.ListBranches()
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: branches}, nil
}

func (c *controller) GetBases(ctx iris.Context) (*controllers.GenericResponse, error) {
	branches, err := c.branchService.ListMilateryBases()
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: branches}, nil
}

func (c *controller) GetBy(ctx iris.Context, id int64) (*controllers.GenericResponse, error) {
	branch, err := c.branchService.GetBranch(id)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: branch}, nil
}

func (c *controller) GetLocationByBranches(ctx iris.Context, locationID int64) (*controllers.GenericResponse, error) {
	branches, err := c.branchService.GetBranchesByLocation(locationID)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: branches}, nil
}

func (c *controller) Post(ctx iris.Context) (*controllers.GenericResponse, error) {
	branch := &models.MBranch{}
	err := ctx.ReadJSON(branch)
	if err != nil {
		return nil, err
	}
	if len(branch.Title) < 2 {
		return nil, errors.New("عنوان باید بیشتر از 2 کاراکتر باشد")
	}
	branch.Alias = branch.Alias + " " + branch.Title
	branch.CreatedAt = time.Now()
	branch.UpdatedAt = time.Now()
	// var RootBranch *models.MBranch
	if branch.ID != 0 {
		rootBranch, err := c.branchService.GetBranch(branch.ID)
		if err != nil {
			return nil, err
		}
		if rootBranch.Xid != 0 {
			return nil, errors.New("شما نمیتوانید برای یگان قسمت تعریف کنید")
		}
		// RootBranch = rootBranch
	}
	if branch.LocationID != 0 && branch.Xid == 0 {
		return nil, errors.New("قسمت نمیتواند شهر داشته باشد")
	}
	if branch.Xid == 1 && branch.LocationID == 0 {
		return nil, errors.New("یگان ها باید شهر داشته باشند")
	}
	if branch.Xid == 1 {
		if branch.NativeServiceDurationMonth == 0 || branch.NonNativeServiceDurationMonth == 0 {
			return nil, errors.New("مدت خدمت بومی و غیر بومی باید مشخص شود")
		}
	}

	err = c.branchService.CreateBranch(branch)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: branch}, nil
}

func (c *controller) Put(ctx iris.Context) (*controllers.GenericResponse, error) {
	branch := &models.MBranch{}
	err := ctx.ReadJSON(branch)
	if err != nil {
		return nil, err
	}
	if len(branch.Title) < 2 {
		return nil, errors.New("عنوان باید بیشتر از 2 کاراکتر باشد")
	}
	branch.Alias = branch.Alias + " " + branch.Title
	branch.UpdatedAt = time.Now()
	// var RootBranch *models.MBranch
	if branch.ID != 0 {
		rootBranch, err := c.branchService.GetBranch(branch.ID)
		if err != nil {
			return nil, err
		}
		if rootBranch.Xid != 0 {
			return nil, errors.New("شما نمیتوانید برای یگان قسمت تعریف کنید")
		}
		// RootBranch = rootBranch
	}
	if branch.LocationID != 0 && branch.Xid == 0 {
		return nil, errors.New("قسمت نمیتواند شهر داشته باشد")
	}
	if branch.Xid == 1 && branch.LocationID == 0 {
		return nil, errors.New("یگان ها باید شهر داشته باشند")
	}
	if branch.Xid == 1 {
		if branch.NativeServiceDurationMonth == 0 || branch.NonNativeServiceDurationMonth == 0 {
			return nil, errors.New("مدت خدمت بومی و غیر بومی باید مشخص شود")
		}
	}
	err = c.branchService.UpdateBranch(branch)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true, Data: branch}, nil
}

func (c *controller) DeleteBy(ctx iris.Context, id int64) (*controllers.GenericResponse, error) {
	token := ctx.GetHeader("Authorization")
	_, err := c.userService.Acl(token, "branch", models.PermissionTypeDelete)
	if err != nil {
		return nil, err
	}
	err = c.branchService.DeleteBranch(id)
	if err != nil {
		return nil, err
	}
	return &controllers.GenericResponse{Success: true}, nil
}

// func (c *controller) Get
func (c *controller) Options(ctx iris.Context) {
	ctx.StatusCode(200)
}

func (c *controller) OptionsBy(ctx iris.Context) {
	ctx.StatusCode(200)
}
