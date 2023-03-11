package delivery

import (
	"immersiveApp/app/middlewares"
	"immersiveApp/features/auth"
	"immersiveApp/features/users/delivery"
	"immersiveApp/utils/helpers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	Service auth.AuthServiceInterface
}

func New(s auth.AuthServiceInterface) *AuthHandler {
	return &AuthHandler{
		Service: s,
	}
}

func (u *AuthHandler) Login(c echo.Context) error {
	loginRequest := LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}

	token, err := u.Service.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, helpers.ResponseFail("User not found"))
	}

	tokesResponse := map[string]any{
		"token": token,
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Login Success", tokesResponse))
}

func (u *AuthHandler) Register(c echo.Context) error {
	registerRequest := delivery.UserRequest{}
	if err := c.Bind(&registerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("error bind data"))
	}
	user := delivery.UserRequestToUserEntity(registerRequest)

	if err := u.Service.Register(user); err != nil {
		return c.JSON(http.StatusInternalServerError, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", registerRequest))
}

func (u *AuthHandler) GetUserLogin(c echo.Context) error {
	tokenClaim := middlewares.ClaimsToken(c)
	return c.JSON(http.StatusOK, helpers.ResponseSuccess("-", tokenClaim))
}

func (u *AuthHandler) ChangePassword(c echo.Context) error {
	user_id := middlewares.ClaimsToken(c).Id

	r := ChangePasswordRequest{}
	if err := c.Bind(&r); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail("Error bind data"))
	}

	if err := u.Service.ChangePassword(uint(user_id), r.OldPassword, r.NewPassword, r.ConfirmPassword); err != nil {
		return c.JSON(http.StatusBadRequest, helpers.ResponseFail(err.Error()))
	}

	return c.JSON(http.StatusOK, helpers.ResponseSuccess("Change password Success", nil))
}
