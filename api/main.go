package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/config"
	"api/infrastructure"
	"api/interface/handler"
	"api/usecase"
)

func init() {
	config.LoadEnv()
}

func main() {
	// repository
	userRepository := infrastructure.NewUserRepository()
	clubRepository := infrastructure.NewClubRepository()
	// usecase
	userUsecase := usecase.NewUserUsecase(userRepository)
	clubUsecase := usecase.NewClubUsecase(clubRepository)
	// handler
	userHandler := handler.NewUserHandler(userUsecase)
	clubHandler := handler.NewClubHandler(clubUsecase)

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{"*"},
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))

	// Router
	handler.InitRouting(
		e,
		userHandler,
		clubHandler,
	)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.GetEnv("PORT", "8080")))
}
