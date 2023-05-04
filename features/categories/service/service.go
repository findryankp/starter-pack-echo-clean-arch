package service

import (
	"cleanarc/features/categories"

	"github.com/go-playground/validator/v10"
)

type Service struct {
	Data     categories.DataInterface
	validate *validator.Validate
}

func New(data categories.DataInterface) categories.ServiceInterface {
	return &Service{
		Data:     data,
		validate: validator.New(),
	}
}

func (s *Service) GetAll() ([]categories.Core, error) {
	return s.Data.SelectAll()
}

func (s *Service) GetById(id uint) (categories.Core, error) {
	return s.Data.SelectById(id)
}

func (s *Service) Create(dataCore categories.Core) (categories.Core, error) {
	s.validate = validator.New()
	if err := s.validate.Struct(dataCore); err != nil {
		return categories.Core{}, err
	}

	id, err := s.Data.Store(dataCore)
	if err != nil {
		return categories.Core{}, err
	}

	return s.Data.SelectById(id)
}

func (s *Service) Update(dataCore categories.Core, id uint) (categories.Core, error) {
	if checkDataExist, err := s.Data.SelectById(id); err != nil {
		return checkDataExist, err
	}

	s.validate = validator.New()
	if err := s.validate.Struct(dataCore); err != nil {
		return categories.Core{}, err
	}

	err := s.Data.Edit(dataCore, id)
	if err != nil {
		return categories.Core{}, err
	}
	return s.Data.SelectById(id)
}

func (s *Service) Delete(id uint) error {
	if _, err := s.Data.SelectById(id); err != nil {
		return err
	}

	return s.Data.Destroy(id)
}
