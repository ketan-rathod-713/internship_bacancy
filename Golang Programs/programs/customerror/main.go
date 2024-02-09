package main

import (
	"errors"
	"fmt"
)

var SERVER_ERROR = NewCustomError("SERVER_ERROR")
var DATABASE_ERROR = errors.New("DATABASE_ERROR")

// CustomError is a struct representing a custom error type.
type CustomError struct {
	message string
}

// NewCustomError creates and returns a pointer to a new CustomError instance.
func NewCustomError(msg string) error {
	return &CustomError{message: msg}
}

// Error returns the error message.
func (e *CustomError) Error() string {
	return e.message
}

func Process() error {
	return SERVER_ERROR
}

func checkDatabase() error {
	return DATABASE_ERROR
}

func main() {
	// Create a new custom error
	err := NewCustomError("SERVER_ERROR")
	err3 := NewCustomError("SERVER_ERROR")

	err4 := errors.New("DATABASE_ERROR")
	err5 := errors.New("DATABASE_ERROR")

	// Use the custom error
	fmt.Println("Error:", err.Error())
	fmt.Println("Error:", err3.Error())

	fmt.Println("Error:", err4.Error())
	fmt.Println("Error:", err5.Error())

	fmt.Println(errors.Is(err, err3))
	fmt.Println(errors.Is(err4, err5))

	fmt.Println(errors.Is(SERVER_ERROR, err))

	// Niche vale dono same hi error ko point kar rahe he jo ki ban chuki he.
	servererr := Process()
	if servererr != nil {
		fmt.Println(errors.Is(SERVER_ERROR, servererr))
	}

	dberr := checkDatabase()
	if dberr != nil {
		fmt.Println(errors.Is(dberr, DATABASE_ERROR))
	}

}
