package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"api/config"
)

// InitRouting routesの初期化
func InitRouting(
	e *echo.Echo,
	userHandler UserHandler,
) {
	// 以下のルーティングはJWT認証が必要
	r := e.Group("")
	r.Use(middleware.JWTWithConfig(*config.JWTConfig()))

	// user関連
	e.GET("/me", userHandler.FindMe())
}
