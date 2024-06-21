package studentservice

import (
	"errors"
	"fmt"
	"schoolApi/models"

	"gorm.io/gorm"
)

type Service interface {
	GetAllStudents() ([](models.Student), error)
	CreateStudent(student *models.Student) (*models.Student, error)
	GetStudentById(id uint64) (*models.Student, error)
	UpdateStudent(student *models.Student) (*models.Student, error)
	DeleteStudent(id uint64) error
	AddParent(studentId uint64, parentId uint64) error
}

type service struct { // implement Service interface
	DB *gorm.DB
}

func New(db *gorm.DB) Service {
	return &service{DB: db}
}

func (s *service) GetAllStudents() ([]models.Student, error) {
	var students []models.Student
	result := s.DB.Preload("Sport").Preload("Parents").Find(&students)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("can not create student") // TODO: get more info about error
	}

	return students, nil
}

func (s *service) CreateStudent(student *models.Student) (*models.Student, error) {
	result := s.DB.Create(student)
	if result.Error != nil || result.RowsAffected == 0 {
		return nil, errors.New("can not create student") // TODO: get more info about error
	}

	return student, nil
}

func (s *service) GetStudentById(id uint64) (*models.Student, error) {
	return nil, nil
}

func (s *service) UpdateStudent(student *models.Student) (*models.Student, error) {
	return nil, nil
}

func (s *service) DeleteStudent(id uint64) error {
	return nil
}

func (s *service) AddParent(studentId uint64, parentId uint64) error {
	// How to make table name dynamic in this case as it will cause to an error if in future the other tables names changes then it will break my code.
	// table name = student_parent
	query := fmt.Sprintf("insert into gormschoolproject.student_parent(student_id, parent_id) values(%v, %v);", studentId, parentId)
	result := s.DB.Exec(query)

	if result.Error != nil {
		return errors.New("error adding parent to student")
	}

	return nil
}
