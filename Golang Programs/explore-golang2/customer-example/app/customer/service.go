package customer

import (
	"postgres-crud/app"
	"postgres-crud/model"

	"gorm.io/gorm"
)

// For customer we will provide following services
type Service interface {
	CreateCustomer(customer model.Customer) error
	GetAllCustomer() ([]model.Customer, error)
	GetOneCustomer(id string) (model.Customer, error)
	UpdateOneCustomer(customer model.Customer, id string) error
	DeleteOneCustomer(id string) error
}

// ham iss type se harek interface method ko implement karege
type service struct {
	DB *gorm.DB
}

func NewService(app *app.App) Service {
	svc := &service{
		DB: app.DB,
	}
	return svc
}

func (s *service) CreateCustomer(customer model.Customer) error {
	return createCustomer(s.DB, customer)
}

func (s *service) GetAllCustomer() ([]model.Customer, error) {
	return getAllCustomer(s.DB)
}

func (s *service) GetOneCustomer(id string) (model.Customer, error) {
	return getOneCustomer(s.DB, id)
}

func (s *service) UpdateOneCustomer(customer model.Customer, id string) error {
	return updateOneCustomer(s.DB, customer, id)
}

func (s *service) DeleteOneCustomer(id string) error {
	return deleteOneCustomer(s.DB, id)
}