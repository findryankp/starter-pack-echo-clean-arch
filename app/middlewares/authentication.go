package middlewares

import (
	"immersiveApp/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := ValidateToken(c); err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFail(err.Error()))
		}
		return next(c)
	}
}
