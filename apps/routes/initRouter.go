package routes

import (
	"cleanarc/apps/middlewares"
	"cleanarc/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	middlewares.Cors(e)
	middlewares.BasicLogger(e)
	e.GET("/", index)
}

func index(c echo.Context) error {
	var data = map[string]interface{}{
		"message":       "Welcome to Cleanarch Starter Pack",
		"developmen_by": "Findryankp",
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", data))
}
