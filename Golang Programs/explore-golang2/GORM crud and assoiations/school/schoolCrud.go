package school

import (
	"fmt"

	"gorm.io/gorm"
)

func SchoolStudentBelongsToRelationship(db *gorm.DB) {
	db.Exec("CREATE SCHEMA IF NOT EXISTS gormschool")
	// createAndInsertSchoolData(db) // TODO: comment it after use
	fmt.Println(getSchoolData(db)) // get school data

	// insertStudentData(db) // TODO: comment it after use

	simpleJoinOfStudentSchool(db)

	numberOfstudentsInSchool()
}

func numberOfstudentsInSchool(db *gorm.DB) {
	var count int64

	joinTable := db.Model(&School{}).Select("school.id as school_id", "school.name as school_name", "student.name as student_name", "student.id as student_id").Joins("inner join gormschool.student on student.school_id = school.id").Scan(&rs)

	db.Model(&School{}).Count(&count).Group("school")
}

func simpleJoinOfStudentSchool(db *gorm.DB) {
	type result struct {
		SchoolId    uint64
		SchoolName  string
		StudentId   uint64
		StudentName string
	}
	var rs []result
	db.Model(&School{}).Select("school.id as school_id", "school.name as school_name", "student.name as student_name", "student.id as student_id").Joins("inner join gormschool.student on student.school_id = school.id").Scan(&rs)

	fmt.Println(rs)
}

func insertStudentData(db *gorm.DB) {
	db.AutoMigrate(&Student{})
	db.Create(&Student{Name: "Rohan", SchoolId: 2})
}

func getSchoolData(db *gorm.DB) []School {
	var schools []School
	db.Find(&schools)
	return schools
}

func createAndInsertSchoolData(db *gorm.DB) {
	db.AutoMigrate(&School{})
	db.Create(&School{Name: "Arpan vidya sankull"})
}
