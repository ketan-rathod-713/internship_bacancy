package main

import (
	"encoding/json"
	"log"
)

type S struct {
	Foo string
}

func unmarshal(s any) {
	err := json.Unmarshal([]byte(`{"Foo": "bar"}`), &s)

	if err != nil {
		log.Println("err", err)
	}

	log.Printf("Value and type of s is %v and %T", s, s)
}

func main() {
	var s S
	unmarshal(s)

	unmarshal(&s)
}
