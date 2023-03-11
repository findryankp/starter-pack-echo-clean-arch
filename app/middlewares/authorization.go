package middlewares

import (
	"immersiveApp/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		claims := ClaimsToken(c)
		if claims.Role != "admin" {
			return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("Don't have authorization to access this endpoint"))
		}
		return next(c)
	}
}

func CanAccess(user_id_db, user_id_token int, role_token string) bool {
	flag := false
	if user_id_token == user_id_db || role_token == "admin" {
		flag = true
	}

	return flag
}
