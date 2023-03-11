package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"immersiveApp/app/middlewares"
	_authData "immersiveApp/features/auth/data"
	_authHandler "immersiveApp/features/auth/delivery"
	_authService "immersiveApp/features/auth/service"
)

func AuthRouter(db *gorm.DB, e *echo.Echo) {
	data := _authData.New(db)
	service := _authService.New(data)
	handler := _authHandler.New(service)

	g := e.Group("/auth")
	g.POST("/register", handler.Register)
	g.POST("/login", handler.Login)
	// g.POST("/forget-password", handler.Create)

	g.Use(middlewares.Authentication)
	g.GET("/users", handler.GetUserLogin)
	g.POST("/change-password", handler.ChangePassword)
}
