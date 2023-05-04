package routes

import (
	categoryData "cleanarc/features/categories/data"
	categoryHandler "cleanarc/features/categories/delivery"
	categoryService "cleanarc/features/categories/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CategoryRouter(db *gorm.DB, e *echo.Echo) {
	data := categoryData.New(db)
	service := categoryService.New(data)
	handler := categoryHandler.New(service)

	g := e.Group("/categories")
	g.GET("", handler.GetAll)
	g.GET("/:id", handler.GetById)
	g.POST("", handler.Create)
	g.PUT("/:id", handler.Update)
	g.DELETE("/:id", handler.Delete)
}
