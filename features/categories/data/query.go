package data

import (
	"cleanarc/features/categories"

	"gorm.io/gorm"
)

type query struct {
	db *gorm.DB
}

func New(db *gorm.DB) categories.DataInterface {
	return &query{
		db: db,
	}
}

func (q *query) SelectAll() ([]categories.Core, error) {
	var dataModel []Category
	if err := q.db.Find(&dataModel); err.Error != nil {
		return nil, err.Error
	}
	return ListModelToCore(dataModel), nil
}

func (q *query) SelectById(id uint) (categories.Core, error) {
	var dataModel Category
	if err := q.db.First(&dataModel, id); err.Error != nil {
		return categories.Core{}, err.Error
	}
	return ModelToCore(dataModel), nil
}

func (q *query) Store(dataCore categories.Core) (uint, error) {
	dataModel := CoreToModel(dataCore)
	if err := q.db.Create(&dataModel); err.Error != nil {
		return 0, err.Error
	}
	return dataModel.ID, nil
}

func (q *query) Edit(dataCore categories.Core, id uint) error {
	dataModel := CoreToModel(dataCore)
	if err := q.db.Where("id", id).Updates(&dataModel); err.Error != nil {
		return err.Error
	}
	return nil
}

func (q *query) Destroy(id uint) error {
	var dataModel Category
	if err := q.db.Delete(&dataModel, id); err.Error != nil {
		return err.Error
	}
	return nil
}
