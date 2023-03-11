package data

import (
	"immersiveApp/features/users"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) users.UserDataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]users.UserEntity, error) {
	var users []User
	if err := q.db.Preload("Team").Find(&users); err.Error != nil {
		return nil, err.Error
	}
	return ListUserToUserEntity(users), nil
}

func (q *query) SelectById(id uint) (users.UserEntity, error) {
	var user User
	if err := q.db.Preload("Team").First(&user, id); err.Error != nil {
		return users.UserEntity{}, err.Error
	}
	return UserToUserEntity(user), nil
}

func (q *query) Store(userEntity users.UserEntity) (uint, error) {
	user := UserEntityToUser(userEntity)
	if err := q.db.Create(&user); err.Error != nil {
		return 0, err.Error
	}
	return user.ID, nil
}

func (q *query) Edit(userEntity users.UserEntity, id uint) (uint, error) {
	user := UserEntityToUser(userEntity)
	if err := q.db.Where("id", id).Updates(&user); err.Error != nil {
		return 0, err.Error
	}

	inputForStatus := map[string]interface{}{"status": user.Status}
	if err := q.db.Model(&user).Where("id", id).Updates(inputForStatus); err.Error != nil {
		return 0, err.Error
	}

	return id, nil
}

func (q *query) Destroy(id uint) error {
	var user User
	if err := q.db.Delete(&user, id); err.Error != nil {
		return err.Error
	}
	return nil
}
