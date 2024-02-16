package studentapi

import (
	studentservice "schoolApi/app/studentService"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func New(db *gorm.DB) *api {
	return &api{
		DB:      db,
		Service: studentservice.New(db),
	}
}

type api struct {
	DB      *gorm.DB
	Service studentservice.Service
}

func Routes(parent_router *mux.Router, db *gorm.DB) {
	studentAPI := New(db)
	parent_router.HandleFunc("/", studentAPI.createStudent).Methods("POST")
	parent_router.HandleFunc("/", studentAPI.getStudents).Methods("GET")
	parent_router.HandleFunc("/parent", studentAPI.AddParent).Methods("POST") // {student_id, parent_id}
}
