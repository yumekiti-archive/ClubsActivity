package config

import (
	"github.com/labstack/echo/v4/middleware"
)

func JWTConfig() *middleware.JWTConfig {
	return &middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}
}
