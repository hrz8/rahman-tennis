package main

import (
	"fmt"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/hrz8/rahman-tennis/models"
	"github.com/hrz8/rahman-tennis/shared/config"
	"github.com/hrz8/rahman-tennis/shared/container"
	"github.com/hrz8/rahman-tennis/shared/database"
	"github.com/hrz8/rahman-tennis/shared/lib"

	ContainerRepository "github.com/hrz8/rahman-tennis/domains/container/repository"
	PlayerRepository "github.com/hrz8/rahman-tennis/domains/player/repository"

	ContainerUsecase "github.com/hrz8/rahman-tennis/domains/container/usecase"
	PlayerUsecase "github.com/hrz8/rahman-tennis/domains/player/usecase"

	ContainerService "github.com/hrz8/rahman-tennis/domains/container/service"
	PlayerService "github.com/hrz8/rahman-tennis/domains/player/service"
)

func main() {
	e := echo.New()

	appContainer := container.DefaultContainer()
	appConfig := appContainer.MustGet("shared.config").(config.AppConfigInterface)
	mysql := appContainer.MustGet("shared.mysql").(database.MysqlInterface)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())

	mysqlSess, err := mysql.Connect()
	if err != nil {
		panic(fmt.Sprintf("[ERROR] failed open mysql connection: %s", err.Error()))
	}

	if len(os.Args) > 1 {
		if os.Args[1] == "migrate" {
			mysqlSess.Debug().AutoMigrate(
				&models.Player{},
				&models.Container{},
			)
		}
	}

	e.Validator = &lib.CustomValidator{Validator: validator.New()}
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ac := &lib.AppContext{
				Context:      c,
				MysqlSession: mysqlSess,
			}
			return next(ac)
		}
	})

	playerRepo := PlayerRepository.NewRepository(mysqlSess)
	containerRepo := ContainerRepository.NewRepository(mysqlSess)

	playerUsecase := PlayerUsecase.NewUsecase(playerRepo, containerRepo)
	containerUsecase := ContainerUsecase.NewUsecase(containerRepo)

	PlayerService.InitService(e, playerUsecase)
	ContainerService.InitService(e, containerUsecase, playerUsecase)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", appConfig.GetAppPort())))
}
