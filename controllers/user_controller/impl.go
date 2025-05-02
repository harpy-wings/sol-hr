package user_controller

import (
	"errors"

	"github.com/harpy-wings/sol-hr/controllers"
	"github.com/harpy-wings/sol-hr/pkg/usermanger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type controller struct {
	userService usermanger.IUserManager
}

func New() (controllers.IController, error) {
	c := new(controller)
	if usermanger.Default == nil {
		return nil, errors.New("usermanger is not initialized")
	}
	c.userService = usermanger.Default
	return c, nil
}

func (c *controller) Register(app *iris.Application) error {
	m := mvc.New(app.Party("/api/user"))
	m.Handle(c)
	return nil
}
func (c *controller) init() error {
	if usermanger.Default != nil {
		c.userService = usermanger.Default
	} else {
		var err error
		c.userService, err = usermanger.New()
		if err != nil {
			return err
		}
	}
	return nil
}
