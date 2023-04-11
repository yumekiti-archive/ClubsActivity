package handler

import (
	_ "api/docs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"net/http"

	"api/config"
)

// InitRouting routesの初期化
func InitRouting(
	e *echo.Echo,
	userHandler UserHandler,
	clubHandler ClubHandler,
	activityHandler ActivityHandler,
) {
	// 以下のルーティングはJWT認証が必要
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	// バージョン管理
	v1 := e.Group("/v1")

	// user関連
	v1.GET("/me", userHandler.FindMe())

	users := v1.Group("/users")
	users.POST("", userHandler.Store())

	// club関連
	club := v1.Group("/clubs")
	club.GET("", clubHandler.FindAll())
	club.GET("/:id", clubHandler.FindByID())
	club.GET("/:id/activities", activityHandler.FindByClubID())
	club.GET("/:id/users", userHandler.FindByClubID())

	// swagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.GET("/swagger", func(c echo.Context) error {
		return c.Redirect(http.StatusMovedPermanently, "/swagger/index.html")
	})
}
