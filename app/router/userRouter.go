package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"immersiveApp/app/middlewares"
	_userData "immersiveApp/features/users/data"
	_userHandler "immersiveApp/features/users/delivery"
	_userService "immersiveApp/features/users/service"
)

func UserRouter(db *gorm.DB, e *echo.Echo) {
	data := _userData.New(db)
	service := _userService.New(data)
	handler := _userHandler.New(service)

	g := e.Group("/users")
	g.Use(middlewares.Authentication)
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
