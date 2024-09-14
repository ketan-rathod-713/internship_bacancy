package studentapi

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"schoolApi/models"
)

var apiLogs = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
var errorLogs = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)

func (a *api) getStudents(w http.ResponseWriter, r *http.Request) {
	apiLogs.Println("GET students")

	students, err := a.Service.GetAllStudents()

	if err != nil {
		errorLogs.Println("Error Fetching Json Data")
		// send json eror here
	}

	json.NewEncoder(w).Encode(students)
}

func (a *api) createStudent(w http.ResponseWriter, r *http.Request) {
	apiLogs.Println("GET students")

	var student *models.Student = &models.Student{}
	err := json.NewDecoder(r.Body).Decode(student)

	if err != nil {
		errorLogs.Println("Error Fetching Json Data", err)
		// send json eror here
	}

	student, err = a.Service.CreateStudent(student)
	if err != nil {
		http.Error(w, "error occured", http.DefaultMaxHeaderBytes)
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(student)
}

func (a *api) AddParent(w http.ResponseWriter, r *http.Request) {
	// TODO: get student id and then add parent for i
	studentParentRelation := models.StudentParentRelationData{}
	json.NewDecoder(r.Body).Decode(&studentParentRelation)

	// Now insert data
	log.Println("Insert data")
	err := a.Service.AddParent(studentParentRelation.StudentId, studentParentRelation.ParentId)
	if err != nil {
		http.Error(w, "An error occured", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(studentParentRelation)
}
