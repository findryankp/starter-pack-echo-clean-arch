package service

import (
	"errors"
	"immersiveApp/app/middlewares"
	"immersiveApp/features/auth"
	"immersiveApp/features/users"
	"immersiveApp/utils/helpers"
)

type authService struct {
	data auth.AuthDataInterface
}

func New(u auth.AuthDataInterface) auth.AuthServiceInterface {
	return &authService{
		data: u,
	}
}

func (u *authService) Login(email, password string) (string, error) {
	if email == "" || password == "" {
		return "", errors.New("email and password must be fill")
	}

	user, err := u.data.GetUserByEmailOrId(email, 0)
	if err != nil || !helpers.CheckPasswordHash(password, user.Password) {
		return "", errors.New("user and password not found")
	}

	//make jwt
	token, errToken := middlewares.CreateToken(int(user.Id), user.Role)
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}

func (u *authService) Register(request users.UserEntity) error {
	return u.data.Register(request)
}

func (u *authService) ChangePassword(id uint, oldPassword, newPassword, confirmPssword string) error {
	if oldPassword == "" || newPassword == "" || confirmPssword == "" {
		return errors.New("old password,new password, and confirm password cannot be empty")
	}

	if newPassword != confirmPssword {
		return errors.New("new password and confirm password must be similarity")
	}

	user, err := u.data.GetUserByEmailOrId(".", id)
	if err != nil || !helpers.CheckPasswordHash(oldPassword, user.Password) {
		return errors.New("old password not match with exist password")
	}

	hash, _ := helpers.HashPassword(newPassword)
	return u.data.EditPassword(id, hash)
}
