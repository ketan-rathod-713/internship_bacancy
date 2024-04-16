package example

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Student struct {
	Name string `json:"name"`
}

func Decoder() {
	reader := strings.NewReader(`{"name": "ketan"}`)
	decoder := json.NewDecoder(reader)

	decoder.UseNumber()
	// token, err := decoder.Token()

	var s Student
	err := decoder.Decode(&s)

	// fmt.Println("Token and error are", token, err)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("marshalled struct is ", s)

	fmt.Println("Decoding inside map")
	var mp map[string]interface{}

	// it will return EOF
	err = decoder.Decode(&mp)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(mp)
}

// How decode works internally
