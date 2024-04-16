package main

import (
	"fmt"
	"gormnotes/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitialiseDB(config *models.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%v port=%v user=%v password=%v", config.HOST, config.DB_PORT, config.DB_USER, config.DB_USER_PASSWORD)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	config := models.Config{
		PORT:             ":8080",
		DB_PORT:          "5432",
		DATABASE:         "bacancy",
		HOST:             "localhost",
		DB_USER:          "bacancy",
		DB_SCHEMA_NAME:   "gorm_final_revision",
		DB_USER_PASSWORD: "admin",
	}

	db, err := InitialiseDB(&config)

	if err != nil {
		log.Fatal(err)
	}

	// migrate the schema
	db.AutoMigrate(&models.Student{}, &models.Parent{})

	for {
		showCommands()
		var selected int
		_, err = fmt.Scan(&selected)
		if err != nil {
			log.Fatal(err)
		}

		switch selected {
		case 1:
			fmt.Println("Creating A Parent...")
			fmt.Println("Enter Father name and mother name")

			var fathername string
			var mothername string
			fmt.Scan(&fathername, &mothername)

			parent := models.Parent{
				FatherName: fathername,
				MotherName: mothername,
			}

			result := db.Create(&parent)

			fmt.Println(result.Error)

			fmt.Println("Parents Information Saved to database")
			fmt.Println("Id if parents is ", parent.Id)
			fmt.Println("")
		case 2:
			fmt.Println("Fetching all parents...")
			var parents []models.Parent
			result := db.Find(&parents)

			fmt.Println(result.Error)

			fmt.Println("All Parents")
			fmt.Printf("%v %v %v %v \n", "sr no", "id", "par.FatherName", "par.MotherName")
			for i, par := range parents {
				fmt.Printf("%v %v %v %v \n", i, par.Id, par.FatherName, par.MotherName)
			}
			fmt.Println("")

		case 3:
			fmt.Println("Creating a student")

			fmt.Println("Enter your full name like firname lastname sername")
			var firstname string
			var lastname string
			var sername string
			fmt.Scan(&firstname, &lastname, &sername)

			fmt.Println("give parent id of student")
			var parentId int
			fmt.Scan(&parentId)

			student := models.Student{
				PersonInfo: models.PersonInfo{
					FirstName: firstname,
					LastName:  lastname,
					SerName:   sername,
				},
				ParentRefer: parentId,
			}

			db.Create(&student)
			fmt.Println("")

		case 4:
			fmt.Println("Fetching all students")

			var students []*models.Student
			db.Preload("Parent").Find(&students)

			for _, stud := range students {
				fmt.Println(stud.Id, stud.FirstName, stud.Parent, stud.Cousins)
			}
			fmt.Println("")
		case 5:

		}
	}
}

func showCommands() {
	fmt.Println("select one of the below option")

	fmt.Println("1. Create one parent")
	fmt.Println("2. Get all parents")
	fmt.Println("3. Create a student")
	fmt.Println("4. Get all students with preloaded parents and cousins information.")
}
