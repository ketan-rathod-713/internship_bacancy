package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Employees []Employee `json:"employees"`
}

type Employee struct {
	Name       string `json:"name"`
	Rating     int    `json:"rating"`
	Department string `json:"department"`
	IsGoodProgress bool `json:"isGoodProgress"`
}

type department struct {
	employeesRatingSum int
	employeesCount     int
	average            float64
}

func main() {
	// Problem Definition:
	// You are given a JSON string representing employee data of a company.
	// Each employee has a name, rating, and department.
	// rating value can be 1 to 10.
	// Your task is to write a Go program to process this JSON data.

	// OBJECTIVE 1: You need to design a solution to calculate the average rating for each department and print the final average rating for each department.

	// hold two values mp[Engineering] = {ratingSum, countOfEmployee}

	// OBJECTIVE 2: Identify employees in each department whose rating is above the department's average and print their names.

	// check for employess in one loop and add key as below

	// OBJECTIVE 3: For employees whose rating is above the department's average, add a new key in the employee list named `isGoodProgress` with the value true; for below average add `isGoodProgress` with the value false.

	data := `{
                "employees": [
                    {"name": "Alice", "rating": 5, "department": "Engineering"},
                    {"name": "Bob", "rating": 6, "department": "Engineering"},
                    {"name": "Charlie", "rating": 7, "department": "Engineering"},
                    {"name": "David", "rating": 8, "department": "Marketing"},
                    {"name": "Eve", "rating": 9, "department": "Marketing"}
                ]
            }`

	var v Data
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		panic(err)
	}
	employees := v.Employees

	mp := make(map[string]*department)

	// TODO:NOT REMOVING JUST FOR LEARNING PURPOSE
	// UnAddresssable field assign ERROR // you can not get the address of struct inside the map
	// for _, e := range employees {
	// 	if val, exist := mp[e.Department]; exist {
	// 		mp[e.Department] = department{ // each time i am creating new struct // and not updating underlying. How to do it.
	// 			employeesRatingSum: val.employeesRatingSum + e.Rating,
	// 			employeesCount:     val.employeesCount + 1,
	// 		}
	// 	} else {
	// 		mp[e.Department] = department{
	// 			employeesRatingSum: e.Rating,
	// 			employeesCount:     1,
	// 		}
	// 	}
	// }

	for _, e := range employees {
		if _, exist := mp[e.Department]; exist {
			dp := mp[e.Department]
			dp.employeesCount += 1
			dp.employeesRatingSum += e.Rating
		} else {
			mp[e.Department] = &department{
				employeesRatingSum: e.Rating,
				employeesCount:     1,
			}
		}
	}

	// fmt.Println(*mp["Engineering"])
	// Calculate average for all departments
	fmt.Println("AVERAGE RATING DEPARTMENT WISE")
	for dName, val := range mp {
		val.average = float64(val.employeesRatingSum) / float64(val.employeesCount)
		fmt.Printf("%v - %v \n", dName, val.average)
	}


	fmt.Println("\nEXCEPTIONAL EMPLOYEES")
	// Now iterate over employees and do objectives
	for i, e := range employees {
		if float64(e.Rating) > mp[e.Department].average {
			fmt.Println(e.Name)
			// e.IsGoodProgress = true // why it is not updating // is it not referencing same ??
			employees[i].IsGoodProgress = true
		} else {
			e.IsGoodProgress = false
		}
	}

	fmt.Println("ALL EMPLOYEES")
	fmt.Printf("%v \t %v \t %v \t %v\n", "Name", "Rating","Department", "IsGoodProgress")
	for _, e := range employees {
		fmt.Printf("%v \t %v \t %v \t %v\n", e.Name, e.Rating, e.Department, e.IsGoodProgress)
	}
	
}
