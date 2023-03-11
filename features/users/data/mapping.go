package data

import (
	"immersiveApp/features/users"
)

func UserEntityToUser(userEntity users.UserEntity) User {
	return User{
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		Password:    userEntity.Password,
		PhoneNumber: userEntity.PhoneNumber,
		Address:     userEntity.Address,
		Role:        userEntity.Role,
		Status:      userEntity.Status,
	}
}

func UserToUserEntity(user User) users.UserEntity {

	result := users.UserEntity{
		Id:          user.ID,
		FullName:    user.FullName,
		Email:       user.Email,
		Password:    user.Password,
		PhoneNumber: user.PhoneNumber,
		Address:     user.Address,
		Role:        user.Role,
		Status:      user.Status,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return result
}

func ListUserToUserEntity(user []User) []users.UserEntity {
	var userEntity []users.UserEntity
	for _, v := range user {
		userEntity = append(userEntity, UserToUserEntity(v))
	}
	return userEntity
}
