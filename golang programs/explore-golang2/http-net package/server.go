package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"text/template"

	_ "github.com/lib/pq"
)

type User struct { // if * then it can result in nil pointer reference when using with scan
	Id          string
	FirstName   string
	LastName    string
	Email       string
	Phone       string
	DateOfBirth string
}

type Student struct {
	Name string
	Age  int
}

var db *sql.DB = ConnectDb()

func hello2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path != "/hello2" {
		http.NotFound(w, r)
		return // here we have to return also ha ha
	}

	if r.Method != "POST" { // Only Allowing Post Request on this route.
		http.Error(w, "Method is not supported.", http.StatusNotFound) // 3rd parameter is status code of response writter
		return
	}

	w.Write([]byte("hii ha ha"))
}

func ConnectDb() *sql.DB {
	db, err := sql.Open("postgres", "postgres://bacancy:admin@localhost/bacancy?sslmode=disable")
	CheckError(err)

	// defer db.Close()

	CheckError(db.Ping())

	return db
}

func main() {
	// Using Templates

	temp1 := template.New("template1")

	temp1, _ = temp1.Parse("Hello {{.Name}}, Your Marks are {{.Age}}%!")

	student := Student{
		Name: "Ketan",
		Age:  25,
	}

	err := temp1.Execute(os.Stdout, student)
	CheckError(err)

	// We’ll use the HandleFunc function to add route handlers to the web server.
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) { // Without StriPFunc it will not serve /hello/hello2 or likewise routes.
		w.Write([]byte("Hello, world!"))
	})

	http.HandleFunc("/hello2", hello2Handler)

	// Difference Between Handle and HandleFunc
	// For using Handle we need to define custome type which has implemented handler interface. It should have the ServeHttp Method. // How and when to define one ??
	// HandleFunc is the wrapper around handle func. in which we just need to specify the handler function directly and use it.

	// Home Page
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, world!"))
	// })

	// Why file server on /files not working

	// File Server To host static files.
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/public/", http.StripPrefix("/public/", fs)) //TODO: StripPrefix creates a new handler. ANd how it will see it relative to our folder and price.

	// Form Handler
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/users", usersDataHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)) // Here nil  because we are setting up http2 here hence no need to define it.
}

func usersDataHandler(w http.ResponseWriter, r *http.Request) {
	// CREATE TEMPLATE FILE
	// Get users data

	query := "SELECT *  FROM httpnet.user;"

	rows, err := db.Query(query)
	CheckError(err)

	defer rows.Close()

	users := []User{} // array of users
	for rows.Next() {
		user := User{}
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Phone, &user.DateOfBirth)
		CheckError(err)

		users = append(users, user)
	}

	// fmt.Fprintln(w, "Users Data")

	t, err := template.ParseFiles("users.html")

	t.Execute(w, users)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Only Accept Post Request
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm err: %v", err)
			return
		}

		// For other HTTP methods, or when the Content-Type is not application/x-www-form-urlencoded, the request Body is not read, and r.PostForm is initialized to a non-nil, empty value.
		// Hence Form jab submit hota he tab browser apne aap se form-urlencoded type likh ke send kar deta he data. // see it how it works.

		// Form is parsed now we can get data

		fmt.Println("Post Request Successfull")

		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		dateofbirth := r.FormValue("dateofbirth")
		email := r.FormValue("email")
		phone := r.FormValue("phone")

		fmt.Println("Type of phone is %T", phone)

		// TODO:See form related data parsing and all. How it works ?

		fmt.Println("Got Values")

		fmt.Println(firstname, lastname, dateofbirth, email, phone)

		// Now connected hence submit form // yahi pe stop ho jao
		err := submitForm(firstname, lastname, dateofbirth, email, phone)

		if err != nil {
			http.Error(w, "Invalid Entry", http.StatusBadRequest)
		} else {
			fmt.Fprintf(w, "Form Submitted")
		}
	} else {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func submitForm(firstname string, lastname string, dateofbirth string, email string, phone string) error {

	// connection to he hi ha ha
	query := `INSERT INTO httpnet.user(firstname, lastname, dateofbirth, email, phone) VALUES($1, $2, $3, $4, $5);`

	// start transaction here TODO:

	_, err := db.Exec(query, firstname, lastname, dateofbirth, email, phone)
	if err != nil {
		return err
	}

	return nil
}
