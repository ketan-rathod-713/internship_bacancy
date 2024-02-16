package sportapi

import (
	sportservice "schoolApi/app/sportService"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *api {
	return &api{
		DB:      db,
		Service: sportservice.New(db),
	}
}

type api struct {
	DB      *gorm.DB
	Service sportservice.Service
}

func Routes(parent_router *mux.Router, db *gorm.DB) {
	sportAPI := New(db)
	parent_router.HandleFunc("/", sportAPI.createSport).Methods("POST")
	parent_router.HandleFunc("/", sportAPI.getAllSports).Methods("GET")
}
