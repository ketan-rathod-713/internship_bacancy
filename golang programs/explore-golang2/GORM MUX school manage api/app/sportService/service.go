package sportservice

import (
	"errors"
	"schoolApi/models"

	"gorm.io/gorm"
)

type Service interface {
	GetAllSports() ([]*models.Sport, error)
	CreateSport(sport *models.Sport) (*models.Sport, error)
	DeleteSport(id uint64) error
}

type service struct { // implement Service interface
	DB *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{DB: db}
}

func (s *service) GetAllSports() ([]*models.Sport, error) {
	return nil, nil
}

func (s *service) CreateSport(sport *models.Sport) (*models.Sport, error) {
	result := s.DB.Create(sport)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("can not create student") // TODO: get more info about error
	}

	return sport, nil
}

func (s *service) DeleteSport(id uint64) error {
	return nil
}
