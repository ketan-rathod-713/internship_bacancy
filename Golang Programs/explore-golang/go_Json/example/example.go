package example

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type About struct {
	Id    int    `json:"id"`
	Color string `json:"color"`
	Age   int    `json:"age`
}

func (a *About) UnmarshalJSON(data []byte) error {
	// first unmarshal it using map
	fmt.Println("what i got", a)

	var mp map[string]interface{}
	err := json.Unmarshal(data, &mp)

	fmt.Println("map", mp)

	if a == nil {
		fmt.Println("nil")
		return nil
	}

	color := "white"
	if val, ok := mp["colorId"]; ok {
		if val == float64(1) {
			color = "red"
		} else if val == float64(2) {
			color = "grey"
		}
	}

	a.Color = color

	// now set the data according to our rule
	id, ok := mp["id"]
	if ok {
		idInt, ok := id.(float64)
		if ok {
			a.Id = int(idInt)
		} else {
			fmt.Println("id is not a int", id)
		}
	}

	age, ok := mp["age"]
	if ok {
		ageInt, ok := age.(float64)
		if ok {
			a.Age = int(ageInt)
		} else {
			fmt.Println("age is not int", age)
		}
	}

	// may be above condtions not run

	// by default float64 value aave che
	// a.Id = int(id.(float64))
	// a.Age =

	if err != nil {
		return err
	}

	return nil
}

// we can also use map inside it and call for marshal json function for it. but don't call marshal json function for this struct as it will go in infinite loop ha ha
func (a *About) MarshalJSON() ([]byte, error) {
	// i don't want to marshal it normally.

	fmt.Println("what i got", a)

	if a.Id != 1 {
		return []byte("{}"), nil
	}

	// means
	fmt.Println("Marshaller called for about")
	custom := fmt.Sprintf(`{"about_id":%v, "about_color": "%s", "about_age": %v}`, a.Id, a.Color, a.Age)
	fmt.Println(custom)
	return []byte(custom), nil

	// NOTE: calling here marshal json will go in infinite loop ha ha
}

type animal struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	About *About `json:"about,omitempty"`
}

// Stringer interface
// we should write stringer interface carefully as some pointer values can be null too.
func (a animal) String() string {
	if a.About != nil {
		return fmt.Sprintf("name %v, type %v \n about id %v age %v color %v", a.Name, a.Type, a.About.Id, a.About.Age, a.About.Color)
	} else {
		return fmt.Sprintf("name %v type %v about is null", a.Name, a.Type)
	}
}

// func (a animal) MarshalJSON() ([]byte, error) {
// 	return , nil
// }

func Example() {
	fmt.Println("Learning encoding/json")

	a := animal{
		Name: "good <> boy",
		Type: "buffello",
		About: &About{
			Id:    1,
			Color: "red",
			Age:   30,
		},
	}

	fmt.Println(a)

	// By default HtmlEscape is applied so that < > inside string is converted to $lt or $gt etc. runes.
	data, err := json.Marshal(a)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))

	var out bytes.Buffer
	json.Indent(&out, data, "", "\t")

	out.WriteTo(os.Stdout)

	fmt.Println("")

	// Unmarshalling

	// str := `{"name": "ketan rathod", "type": "human", "about": {"id": 1, "colorId": 2}}`
	str := `{"name":"ketan rathod"}`

	// using map // map of map in case of nested structures.
	var mp map[string]interface{}
	err = json.Unmarshal([]byte(str), &mp)

	if err != nil {
		fmt.Println("error occured", err)
	}

	fmt.Println(mp)
	fmt.Println(mp["name"])

	// now using structure and with custom logic

	fmt.Println("Now using struct")
	var k animal
	err = json.Unmarshal([]byte(str), &k)

	if err != nil {
		fmt.Println("error occured", err)
		return
	}

	fmt.Println(k)
}

// How to not return some values and mark it as null for ex.

// Remarks

/*
[]byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON value.
The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value,
defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.

As a special case, if the field tag is "-", the field is always omitted. Note that a field with name "-"
can still be generated using the tag "-,".

View how all data types such as Map, Slice, Interface, Struct and Channel are stored internally and how it works.
*/
