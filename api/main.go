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
	// database
	db := config.NewDB()

	// repository
	userRepository := infrastructure.NewUserRepository(db)
	clubRepository := infrastructure.NewClubRepository(db)
	activityRepository := infrastructure.NewActivityRepository(db)
	// usecase
	userUsecase := usecase.NewUserUsecase(userRepository)
	clubUsecase := usecase.NewClubUsecase(clubRepository)
	activityUsecase := usecase.NewActivityUsecase(activityRepository)
	// handler
	userHandler := handler.NewUserHandler(userUsecase)
	clubHandler := handler.NewClubHandler(clubUsecase)
	activityHandler := handler.NewActivityHandler(activityUsecase)

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
		activityHandler,
	)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.GetEnv("PORT", "8080")))
}
