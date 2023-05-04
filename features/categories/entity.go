package categories

import "time"

type Core struct {
	Id        uint
	Category  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll() ([]Core, error)
	GetById(id uint) (Core, error)
	Create(data Core) (Core, error)
	Update(data Core, id uint) (Core, error)
	Delete(id uint) error
}

type DataInterface interface {
	SelectAll() ([]Core, error)
	SelectById(id uint) (Core, error)
	Store(data Core) (uint, error)
	Edit(data Core, id uint) error
	Destroy(id uint) error
}
