package parentservice

import (
	"errors"
	"schoolApi/models"

	"gorm.io/gorm"
)

type Service interface {
	CreateParent(parent *models.Parent) (*models.Parent, error)
	GetAllParents() ([]*models.Parent, error)
}

type service struct { // implement Service interface
	DB *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{DB: db}
}

func (s *service) GetAllParents() ([]*models.Parent, error) {
	return nil, nil
}

func (s *service) CreateParent(parent *models.Parent) (*models.Parent, error) {
	result := s.DB.Create(parent)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("can not create student") // TODO: get more info about error
	}

	return parent, nil
}
