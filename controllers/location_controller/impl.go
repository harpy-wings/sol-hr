package location_controller

import (
	"errors"

	"github.com/harpy-wings/sol-hr/controllers"
	"github.com/harpy-wings/sol-hr/pkg/geoManager"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type controller struct {
	geoService geoManager.GeoManager
}

func New() (controllers.IController, error) {
	c := new(controller)
	if geoManager.Default == nil {
		return nil, errors.New("geoManager is not initialized")
	}
	c.geoService = geoManager.Default
	// c.setDefaults()
	// for _, v := range ops {
	// 	err := v(c)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }
	// err := c.init()
	// if err != nil {
	// 	return nil, err
	// }
	return c, nil
}

func (c *controller) Register(App *iris.Application) error {
	m := mvc.New(App.Party("/api/locations"))
	m.Handle(c)
	return nil
}
