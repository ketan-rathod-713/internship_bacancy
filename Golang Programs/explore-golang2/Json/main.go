package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

var jsonData = `{
	"name": "John Doe",
	"age": 30,
	"isStudent": false,
	"courses": ["math", "history", "chemistry"]
  }`

type Person struct {
	name      string
	age       int
	isStudent bool
	courses   []string
}

type Person2 struct {
	Name      string
	Age       int
	IsStudent bool
	Courses   []string
}

func main() {
	// marshalling
	// data structure into json

	p := Person{"ketan", 12, true, []string{"good", "one"}}
	jsn, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsn))

	// marshalling for person 2
	p2 := Person2{"ketan", 12, true, []string{"good", "one"}}
	json2, err2 := json.Marshal(p2)
	if err2 != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(json2))

	// Unmarshal
	var p3 Person2
	err3 := json.Unmarshal([]byte(jsonData), &p3)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println("unmarshalled data is ", p3)

	// Custom types using marshler and unmarshler
	marshaling()
}

type Email string
type Name string

// For an email type hamne marhal and unmarshal ka logic likh diya he
func (e Email) MarshalJSON() ([]byte, error) {
	return json.Marshal(strings.ToLower(string(e)))
}
func (e Name) MarshalJSON() ([]byte, error) {
	fmt.Println("Before marshalling ", e)
	return json.Marshal(strings.ToUpper(string(e)))
}

// func (e *Email) UnmarshalJSON(data []byte) error {
// 	var s string
// 	err := json.Unmarshal(data, &s)
// 	if err != nil {
// 		return err
// 	}
// 	*e = Email(strings.ToLower(s))
// 	return nil
// }

type Student struct {
	Name  Name  `json:"name"`
	Email Email `json:"email"` // now not using it then what should i do
}

func marshaling() {
	// ds to json
	p := Student{
		Name:  "John DoeSGSGS",
		Email: "JOHN.DOe@Example.com",
	}

	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(string(jsonData))
}
