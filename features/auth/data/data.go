package data

import (
	"immersiveApp/features/auth"
	"immersiveApp/features/users"
	"immersiveApp/features/users/data"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) auth.AuthDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) GetUserByEmailOrId(email string, id uint) (users.UserEntity, error) {
	var user data.User
	if err := q.db.Where("email = ? or id = ?", email, id).First(&user); err.Error != nil {
		return users.UserEntity{}, err.Error
	}

	return data.UserToUserEntity(user), nil
}

func (q *query) Register(request users.UserEntity) error {
	user := data.UserEntityToUser(request)
	if err := q.db.Create(&user); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) EditPassword(id uint, pass string) error {
	var user data.User
	if err := q.db.Model(&user).Where("id", id).Update("password", pass); err.Error != nil {
		return err.Error
	}
	return nil
}
