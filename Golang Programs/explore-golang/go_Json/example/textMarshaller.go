package example

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Custom type representing a point in 2D space
type Point struct {
	X, Y int
}

// Implementing TextMarshaler interface for Point
func (p *Point) MarshalText() ([]byte, error) {
	// Custom text representation: "X,Y"
	text := fmt.Sprintf("%d,%d", p.X, p.Y)
	return []byte(text), nil
}

// Implementing TextMarshaler interface for Point
func (p *Point) UnmarshalText(data []byte) error {
	// i have data convert into p
	str := string(data)

	slice := strings.Split(str, ",")

	if len(slice) != 2 {
		return errors.New("Give specified formate of point")
	}

	p.X, _ = strconv.Atoi(slice[0])
	p.Y, _ = strconv.Atoi(slice[1])

	return nil
}

func Example2() {
	// Create an instance of Point
	p := Point{X: 10, Y: 20}

	// Marshal Point to text using the custom MarshalText method
	text, err := p.MarshalText()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the marshaled text representation
	fmt.Println("Text representation:", string(text))

	// You can also use encoding/json with TextMarshaler
	data, err := json.Marshal(p)
	if err != nil {
		fmt.Println("JSON Error:", err)
		return
	}

	// Print the JSON representation (which uses the custom text encoding)
	fmt.Println("JSON representation:", string(data))

	type Something struct {
		Point Point
	}

	s := Something{
		Point: Point{
			X: 60,
			Y: 20,
		},
	}

	data, err = json.Marshal(s)
	if err != nil {
		fmt.Println("err", err)
		return
	}

	fmt.Println("Marshalled String ", string(data))

	// let me do unmarshal on Something

	str := `{"Point": "12,30"}`
	// var mp map[string]interface{}
	var something Something
	err = json.Unmarshal([]byte(str), &something)

	if err != nil {
		log.Fatal(err)
	}

	// it will unmarshal that particular text to original struct and similarly struct to that string
	fmt.Println(something)

	// ok ok marshal text is different from the marshal json which is default of marshal function
	// lets try marshalText method

	// json.MarshalText()

	// Ohkk now understood so map's key should implement the TextMarhsaller interface because it can not contain the object values as key inside it.
}
