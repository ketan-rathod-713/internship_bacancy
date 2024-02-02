package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const API_KEY = "fde8a92a6d889973f2366762d5aa8b86"

type Wheather struct {
	Coord struct {
		Lat float64 `json:"lat"`
		Lon float64 `json:"lon"`
	}
	Name string `json:"name"`
	Humidity string `json:"humidity"`
	Main struct {
		Temp float64 `json:"temp"`
        Pressure float64 `json:"pressure"`
		Humidity float32 `json:"humidity"`
	}
	Wind struct {
        Speed float64 `json:"speed"`
        Deg float64 `json:"deg"`
    }
	Sys struct {
        Country string `json:"country"`
        Sunrise int `json:"sunrise"`
        Sunset int `json:"sunset"`
    }
}

func main() {
	var lattitude float64 = 22
	var longitude float64 = 73

	if(len(os.Args) == 3){
		lattitude, _ = strconv.ParseFloat(os.Args[1],10)
		longitude, _ = strconv.ParseFloat(os.Args[2], 10)
		fmt.Println(longitude, lattitude)
	}

	URL := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?lat=%v&lon=%v&appid=%v", lattitude, longitude, API_KEY)

	response, err := http.Get(URL)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	contentBytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println(string(contentBytes))

	var wheather Wheather
	err = json.Unmarshal(contentBytes, &wheather)
	if err != nil {
		panic(err)
	}
	// fmt.Println(wheather)

	fmt.Printf("================================================\n")
	fmt.Println("Today's Wheather Is ")
	fmt.Printf("================================================\n")
	fmt.Printf("Lattitude : %v , Longitude : %v \n", wheather.Coord.Lat, wheather.Coord.Lon)
	fmt.Printf("Name : %v, Country : %v \n", wheather.Name, wheather.Sys.Country)
	fmt.Printf("Humidity : %v \n", wheather.Humidity)
	fmt.Printf("Pressure : %v, Temperature : %v \n", wheather.Main.Pressure, wheather.Main.Temp)
}

// move exe file to usr/../bin/ and then run with the name of the foldername
// this is working project
// run it using by typing wheather in cli with lattitude and longitude information
// can it get my current location and give relvent infromation/
