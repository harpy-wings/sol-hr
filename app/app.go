package app

import (
	"context"
	"errors"
	"fmt"

	// _ "github.com/harpy-wings/solder-hr-fin/docs"

	"github.com/harpy-wings/sol-hr/controllers/branch_controller"
	"github.com/harpy-wings/sol-hr/controllers/location_controller"
	"github.com/harpy-wings/sol-hr/controllers/solder_controller"
	"github.com/harpy-wings/sol-hr/controllers/user_controller"
	"github.com/harpy-wings/sol-hr/models"
	"github.com/harpy-wings/sol-hr/pkg/branchManager"
	"github.com/harpy-wings/sol-hr/pkg/geoManager"
	solderManager "github.com/harpy-wings/sol-hr/pkg/soldermanager"
	"github.com/harpy-wings/sol-hr/pkg/usermanger"
	utilitymanager "github.com/harpy-wings/sol-hr/pkg/utilityManager"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 185.204.171.83:7080
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @description Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345".

type IApp interface {
	Run(context.Context) error
	GracefulStop(context.Context) error
}
type App struct {
	logger     logrus.FieldLogger
	db         *gorm.DB
	iapp       *iris.Application
	onShutdown []func()
	config     struct {
		port int
	}
}

func (a *App) preSet(v *viper.Viper) error {
	logger := logrus.StandardLogger()
	logger.SetLevel(logrus.DebugLevel)
	logger.ReportCaller = true
	a.logger = logger

	{
		// ─── Other ───────────────────────────────────────────────────────────
		p := v.GetInt("port")
		if p == 0 {
			p = 3000
		}
		a.config.port = p
	}

	// db config
	{
		dbraw := v.GetString("db.raw")
		if dbraw == "" {
			a.logger.Warnf("db.raw is not set")
			return errors.New("db.raw is not set")

		}
		dbclient := v.GetString("db.client")
		if dbclient == "" {
			a.logger.Warnf("db.client is not set")
			return errors.New("db.client is not set")
		}
		db, err := gorm.Open(dbFoF(dbclient)(dbraw), &gorm.Config{})
		if err != nil {
			a.logger.Warnf(err.Error())
		}
		a.db = db
		models.DB = a.db
		err = models.InitDB()
		if err != nil {
			a.logger.Warnf(err.Error())

			return err
		}

	}
	// iris app config
	{
		iapp := iris.New()
		if logLvl := v.GetString("logger.level"); logLvl != "" {
			iapp.Logger().SetLevel(logLvl)
		} else {
			iapp.Logger().SetLevel("debug")
		}
		{
			crs := cors.New(cors.Options{
				AllowedOrigins:   []string{"*"},
				AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
				AllowedHeaders:   []string{"*"},
				AllowCredentials: true,
			})

			iapp.Use(crs)

		}
		iapp.Options("/{any:path}", func(ctx iris.Context) { ctx.StatusCode(200) })

		// {
		// 	// swagger
		// 	swaggerUI := swagger.Handler(swaggerFiles.Handler,
		// 		swagger.URL("/swagger/doc.json"),
		// 		swagger.DeepLinking(true),
		// 		swagger.Prefix("/swagger"),
		// 		swagger.SetTheme(swagger.Monokai),
		// 	)
		// 	iapp.Get("/swagger", swaggerUI)
		// 	iapp.Get("/swagger/{any:path}", swaggerUI)
		// 	// Serve the Swagger JSON
		// 	iapp.Get("/swagger/doc.json", func(ctx iris.Context) {
		// 		ctx.ServeFile("./docs/swagger.json")
		// 	})
		// }
		// Server the static files to /
		iapp.Get("/", func(ctx iris.Context) {
			ctx.ServeFile("./static/index.html")
		})
		iapp.Get("/{any:path}", func(ctx iris.Context) {
			ctx.ServeFile("./static/index.html")
		})

		// Serve static assets directory
		iapp.HandleDir("/assets", "./static/assets")
		iapp.HandleDir("/templates", "./templates")
		iapp.HandleDir("/fonts", "./static/fonts")
		iapp.HandleDir("/img", "./static/img")
		a.iapp = iapp
	}

	return nil
}

func (a *App) init(v *viper.Viper) error {
	// os.Setenv(constants.ENV_DRY_RUN, "true")
	{
		utilityManager, err := utilitymanager.New(utilitymanager.WithDB(a.db))
		if err != nil {
			return err
		}
		utilitymanager.Default = utilityManager

		userManager, err := usermanger.New(usermanger.WithDB(a.db))
		if err != nil {
			return err
		}
		usermanger.Default = userManager

		// geoManager
		GeoManager, err := geoManager.New(geoManager.WithDB(a.db))
		if err != nil {
			return err
		}
		geoManager.Default = GeoManager

		// branchmanager
		BranchManager, err := branchManager.New(branchManager.WithDB(a.db), branchManager.WithGeoManager(GeoManager))
		if err != nil {
			return err
		}
		branchManager.Default = BranchManager

		// soldermanager
		SolderManager, err := solderManager.New(solderManager.WithDB(a.db), solderManager.WithGeoManager(GeoManager), solderManager.WithBranchManager(BranchManager), solderManager.WithUtilityManager(utilityManager))
		if err != nil {
			return err
		}
		solderManager.Default = SolderManager
	}
	{
		userController, err := user_controller.New()
		if err != nil {
			return err
		}
		err = userController.Register(a.iapp)
		if err != nil {
			return err
		}

		locationController, err := location_controller.New()
		if err != nil {
			return err
		}
		err = locationController.Register(a.iapp)
		if err != nil {
			return err
		}

		branchController, err := branch_controller.New()
		if err != nil {
			return err
		}
		err = branchController.Register(a.iapp)
		if err != nil {
			return err
		}

		solderController, err := solder_controller.New()
		if err != nil {
			return err
		}
		err = solderController.Register(a.iapp)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) GracefulStop(context.Context) error {
	for _, fn := range a.onShutdown {
		fn()
	}

	return nil
}

func (a *App) Run(ctx context.Context) error {
	return a.iapp.Listen(fmt.Sprintf(":%d", a.config.port))
}

func New(v *viper.Viper) (IApp, error) {
	a := new(App)
	err := a.preSet(v)
	if err != nil {
		return nil, err
	}
	err = a.init(v)
	if err != nil {
		return nil, err
	}
	return a, nil
}

func dbFoF(clientName string) func(raw string) gorm.Dialector {

	switch clientName {
	case "mysql":
		return mysql.Open
	case "postgres":
		return postgres.Open
	case "sqlserver":
		return sqlserver.Open
	case "sqlite":
		return sqlite.Open
	default:
		return mysql.Open
	}
}
