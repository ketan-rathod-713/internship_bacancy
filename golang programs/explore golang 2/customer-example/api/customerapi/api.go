package customerapi

import (
	"postgres-crud/app"
	"postgres-crud/app/customer"
)

// API sidekiq api
type api struct {
	App             *app.App
	CustomerService customer.Service // Type of service interface
}

// New creates a new api
func New(app *app.App) *api {

	return &api{
		App:             app,
		CustomerService: customer.NewService(app), // It is a type of Service interface
	}
}
