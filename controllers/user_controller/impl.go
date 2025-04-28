package user_controller

import (
	"github.com/harpy-wings/sol-hr/pkg/usermanger"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type controller struct {
	userService usermanger.IUserManager
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
