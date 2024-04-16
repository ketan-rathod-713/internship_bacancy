package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("Name %v Age %v", p.Name, p.Age)
}

func (p Student) String() string {
	return fmt.Sprintf("StudentId: %v PersonInfo: %v", p.StudentId, p.Person)
}

type Student struct {
	Person
	StudentId int
}

func (s Student) Learn(){
	fmt.Println("Learn")
}

func (s Student) WorkHard(){
	fmt.Println("Work hard")
}

func (s Student) WorkSmart(){
	fmt.Println("work smart")
}

// can i do interface embedding

type learner interface {
	Learn()
}

type goodLearner interface {
	learner
	WorkSmart()
	WorkHard()
}


func main() {
	var s Student = Student{
		StudentId: 23,
		Person: Person{
			Name: "anad",
			Age: 22,
		},
	}

	fmt.Println(s)

	var l learner = s
	l.Learn()

	var gl goodLearner = s
	gl.Learn()
	gl.WorkHard()
	gl.WorkSmart()
}
