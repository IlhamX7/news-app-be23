package factory

import (
	"news-app-be23/configs"
	userHandler "news-app-be23/internal/features/users/handler"
	userRepository "news-app-be23/internal/features/users/repository"
	userServices "news-app-be23/internal/features/users/services"

	"fmt"
	"news-app-be23/internal/routes"

	"news-app-be23/internal/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitFactory(e *echo.Echo) {
	cfg := configs.ImportSetting()
	db, err := configs.ConnectDB(cfg)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
	}
	if err := db.AutoMigrate(&userRepository.User{}); err != nil {
		fmt.Println("Ada yg bermasalah saat memasukan table user", err.Error())
	}

	pu := utils.NewPasswordUtility()
	jt := utils.NewJwtUtility()

	um := userRepository.NewUserModel(db)
	us := userServices.NewUserService(um, pu, jt)
	uc := userHandler.NewUserController(us)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())

	routes.InitRoute(e, uc)
}
