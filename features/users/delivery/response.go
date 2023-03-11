package delivery

import (
	"immersiveApp/features/users"
)

type UserResponse struct {
	Id          uint   `json:"id,omitempty"`
	TeamId      uint   `json:"team_id,omitempty"`
	FullName    string `json:"full_name"`
	Email       string `json:"email,omitempty"`
	Role        string `json:"role,omitempty"`
	Status      bool   `json:"status"`
	PhoneNumber string `json:"phone_number,omitempty"`
	Address     string `json:"address,omitempty"`
}

func UserEntityToUserResponse(userEntity users.UserEntity) UserResponse {
	result := UserResponse{
		Id:          userEntity.Id,
		FullName:    userEntity.FullName,
		Email:       userEntity.Email,
		Role:        userEntity.Role,
		Status:      userEntity.Status,
		PhoneNumber: userEntity.PhoneNumber,
		Address:     userEntity.Address,
	}

	return result
}

func ListUserEntityToUserResponse(dataCore []users.UserEntity) []UserResponse {
	var dataResponses []UserResponse
	for _, v := range dataCore {
		dataResponses = append(dataResponses, UserEntityToUserResponse(v))
	}
	return dataResponses
}
