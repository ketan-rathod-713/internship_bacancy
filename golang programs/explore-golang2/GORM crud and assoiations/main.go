package main

import (
	"gormcrud/database"
	"gormcrud/school"
	"log"
)

func main() {
	db, err := database.InitialiseDB()
	if err != nil {
		log.Println("error connecting database")
	}

	// // How to set primary key, foregin key and doing all that stuff
	// db.Exec("CREATE SCHEMA IF NOT EXISTS gormbasics1;")

	// db.AutoMigrate(&models.User{}, &models.CreditCard{}, &models.Note{})

	// create one user
	// var user models.User
	// user.ID = 1
	// user.Username = "bacancy"
	// db.Create(&user)

	// Create one note
	// var note models.Note
	// note.UserId = 1
	// db.Create(&note)

	// fetch all notes of given user
	// var note models.Note
	// db.First(&note)
	// it will just give note

	// to get the user corresponding to given note we will do
	// db.Preload("User").First(&note)
	// db.First(&note)

	// var user models.User
	// db.Where("id = ?", note.UserId).Find(&user)

	// log.Println("given note is ", note.ID)
	// log.Println("User for given note is ", note.User.Username)

	// product.ProductPackage(db)

	school.SchoolStudentBelongsToRelationship(db)

}
