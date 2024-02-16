package parentapi

import (
	parentservice "schoolApi/app/parentService"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *api {
	return &api{
		DB:      db,
		Service: parentservice.New(db),
	}
}

type api struct {
	DB      *gorm.DB
	Service parentservice.Service
}

func Routes(parent_router *mux.Router, db *gorm.DB) {
	sportAPI := New(db)
	parent_router.HandleFunc("/", sportAPI.createParent).Methods("POST")
	parent_router.HandleFunc("/", sportAPI.getAllParents).Methods("GET")
}
