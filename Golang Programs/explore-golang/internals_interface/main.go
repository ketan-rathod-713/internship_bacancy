package main

import "fmt"

type about interface {
	getAbout() string
}

type info interface {
	getInfo() string
}

type student struct {
	Id   int
	Name string
}

func (s *student) getAbout() string {
	return fmt.Sprintf("Id: %v, Name: %v", s.Id, s.Name)
}

func (s student) getInfo() string {
	return fmt.Sprintf("Info of student : Name : %v", s.Name)
}

func main() {
	var studentAbout about

	fmt.Println(studentAbout) // it will return nil as it is not pointing to anything

	s := student{
		Name: "ketan rathod",
		Id: 12,
	}

	// In this case s is not implementing this interface but the pointer to the s is implementing it.
	studentAbout = &s

	fmt.Println(studentAbout.getAbout())

	var studentInfo info

	studentInfo = s 
	fmt.Println(studentInfo.getInfo())
}
