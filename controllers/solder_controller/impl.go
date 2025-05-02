package solder_controller

import (
	"errors"

	"github.com/harpy-wings/sol-hr/controllers"
	"github.com/harpy-wings/sol-hr/pkg/branchManager"
	"github.com/harpy-wings/sol-hr/pkg/geoManager"
	solderManager "github.com/harpy-wings/sol-hr/pkg/soldermanager"
	"github.com/harpy-wings/sol-hr/pkg/usermanger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type controller struct {
	geoService    geoManager.GeoManager
	branchService branchManager.IBranchManager
	userService   usermanger.IUserManager
	solderService solderManager.ISolderManager
}

func New() (controllers.IController, error) {
	c := new(controller)
	if geoManager.Default == nil {
		return nil, errors.New("geoManager is not initialized")
	}
	c.geoService = geoManager.Default
	if branchManager.Default == nil {
		return nil, errors.New("branchManager is not initialized")
	}
	c.branchService = branchManager.Default
	if usermanger.Default == nil {
		return nil, errors.New("usermanger is not initialized")
	}
	c.userService = usermanger.Default
	if solderManager.Default == nil {
		return nil, errors.New("solderManager is not initialized")
	}
	c.solderService = solderManager.Default
	// c.setDefaults()
	// for _, v := range ops {
	// 	err := v(c)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// err := c.init()
	// if err != nil {
	// 	return nil, err
	// }
	return c, nil
}

func (c *controller) Register(App *iris.Application) error {
	m := mvc.New(App.Party("/api/solders"))
	m.Handle(c)
	return nil
}
