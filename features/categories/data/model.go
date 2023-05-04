package data

import (
	"cleanarc/features/categories"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category string
}

func CoreToModel(coreData categories.Core) Category {
	return Category{
		Category: coreData.Category,
	}
}

func ModelToCore(modelData Category) categories.Core {
	return categories.Core{
		Id:        modelData.ID,
		Category:  modelData.Category,
		CreatedAt: modelData.CreatedAt,
		UpdatedAt: modelData.UpdatedAt,
	}
}

func ListModelToCore(modelData []Category) []categories.Core {
	var coreData []categories.Core
	for _, v := range modelData {
		coreData = append(coreData, ModelToCore(v))
	}
	return coreData
}
