package router

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func SwaggerRouter(e *echo.Echo) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
