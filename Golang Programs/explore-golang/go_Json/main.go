package main

import (
	"encoding/json"
	"fmt"
)

// lowecase means private
type course struct {
	Name     string `json:"coursename"` // we can create aliases for it here in json ha ha
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // whoever is consuming my api will not get this field ha ha wow great amazing
	Tags     []string `json:"tags,omitempty"` // if anything field is nil then we will not show it wow great.
}

func main() {
	fmt.Println("Welcome to json ha ha")
	EncodeJson()
	EncodeJsonMarshalIndent()
	DecodeJson()
}

// convert data to json // here instead of nil, it places null
func EncodeJson() {
	lcoCourses := []course{
		{"React Js Bootcamp", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Next Js Bootcamp", 499, "lco.in", "abc123", []string{"fullstack-dev", "js"}},
		{"Mern Js Bootcamp", 2299, "lco.in", "abc123", []string{"mern-dev", "js"}},
		{"Angular Js Bootcamp", 199, "lco.in", "abc123", nil},
	}

	// package this data as json data

	// Marshal encodes any interface {} to json
	finalJson, err := json.Marshal(lcoCourses)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

// convert data to json // all the keys are still uppercase // so still we need to do something
func EncodeJsonMarshalIndent() {
	lcoCourses := []course{
		{"React Js Bootcamp", 299, "lco.in", "abc123", []string{"web-dev", "js"}},
		{"Next Js Bootcamp", 499, "lco.in", "abc123", []string{"fullstack-dev", "js"}},
		{"Mern Js Bootcamp", 2299, "lco.in", "abc123", []string{"mern-dev", "js"}},
		{"Angular Js Bootcamp", 199, "lco.in", "abc123", nil},
	}

	// package this data as json data

	// Marshal encodes any interface {} to json
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
        {
                "coursename": "React Js Bootcamp",
                "Price": 299,
                "website": "lco.in",
                "tags": [
                        "web-dev",
                        "js"
                ]
        }
	`)

	// whatever data is coming from web i have structure for it
	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if(checkValid){
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		// we have special syntax for interface to print
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		panic("JSON is not valid")
	}

	// it will take coursname but will see alias and hence vo information will go in Name

	// some cases where you just want to add data to key value pair

	// value is not garranty it can be anything
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	// it is lookig complex but it is not.

	for k, v := range myOnlineData {
		fmt.Printf("Key is %k and value is %v and type is %T \n", k, v, v)
	}
}
