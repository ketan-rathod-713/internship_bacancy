package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // okk now understood as we will not use it and it was giving me error ha ha
)

// var db *sql.DB

// func init() { // Make database connection only once
// 	var err error
// 	connStr := "postgres://bacancy:admin@localhost/bacancy?sslmode=disable"
// 	db, err = sql.Open("postgres", connStr)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if err = db.Ping(); err != nil {
// 		panic(err)
// 	}
// 	// this will be printed in the terminal, confirming the connection to the database
// 	fmt.Println("The database is connected")
// }

type Student struct {
	Id        sql.NullInt32
	Name      sql.NullString
	BirthDate sql.NullString
	Address   sql.NullString
	City      sql.NullString
	Pincode   sql.NullString
	Age       sql.NullInt32
}

const (
	host     = "localhost"
	port     = 5432
	user     = "bacancy"
	password = "admin"
	dbname   = "bacancy"
)

var db *sql.DB

func main() {
	var err error
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected To Database", dbname)

	// Insert unique id
	// insertDataIntoStudent(22,"Aman", "D-34 sarthak row house valthan", "kanpur", "394325", "12-05-2001")

	GetStudents()

}

func GetStudents() {
	rows, err := db.Query(`SELECT id, name, address, city, pincode, birth_date, age FROM school.student;`)
	CheckError(err)

	defer rows.Close()
	for rows.Next() {
		s := Student{}

		err = rows.Scan(&s.Id, &s.Name, &s.Address, &s.City, &s.Pincode, &s.BirthDate, &s.Age)
		CheckError(err)

		fmt.Println(s)
	}
}

func insertDataIntoStudent(id int, name string, address string, city string, pincode string, birthdate string) {
	query := `INSERT INTO school.student(id,name, address, city, pincode, birth_date) 
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := db.Exec(query, id, name, address, city, pincode, birthdate)
	CheckError(err)
}

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXIST report(student_id int, score int);`
	res, err := db.Exec(query)
	CheckError(err)
	fmt.Println(res)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
