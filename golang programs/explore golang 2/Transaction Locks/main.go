package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"
	"time"
)

var mutex sync.Mutex

var count = 0

func main() {
	// var wg sync.WaitGroup // Initialize a WaitGroup

	// Define the book data
	type Student struct {
		Id        uint64 `gorm:"primaryKey" json:"id"`
		Name      string `json:"name"`
		Address   string `json:"address"`
		City      string `json:"city"`
		Pincode   string `json:"pincode"`
		BirthDate string `json:"birth_date"`
		SportId   uint64 `json:"sport_id"` // it will create foreign key here references to Id of Sport
	}

	// Convert the book struct to JSON

	i := 0
	for i < 1000000 {
		p := strconv.Itoa(i)
		jsonData, err := json.Marshal(Student{Name: "aman", SportId: 1, Pincode: p})
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}

		someting(jsonData)
		i++
	}
	time.Sleep(time.Second)
	fmt.Println(count)
}

func someting(jsonData []byte) {
	mutex.Lock()
	count++
	mutex.Unlock()
}
