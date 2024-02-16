package models

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name   string
	Class  int `gorm:"default:18"`
	Age *int `gorm:"default: 5"`
	// School School // association with School
}

// define it in particular schema // by this we can change the default table name
func (s Student) TableName() string {
	return "gormbasics.students"
}

// GORM allows user defined hooks to be implemented for BeforeSave, BeforeCreate, AfterSave, AfterCreate. These hook method will be called when creating a record, refer Hooks for details on the lifecycle

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("THiS IS CALLED BEFORE CREATING STUDENT")
	fmt.Println(s.ID, s.Name, s.Class, s.CreatedAt)

	// if age is greater then 20 then give error ha ha

	if s.Class > 25 {
		return errors.New("ERROR : WRONG CLASS ENTERED")
	}
	return nil
}

// Other types of hooks are BeforeSave, BeforeCreate, AfterSave, AfterCreate TODO: good means if we don;t want to create if wrong data given then give error ha ha // Hence validation can be done at application level.
// We can also skip hooks if we want. // by creating db session with option of {SkipHooks: true}

// type School struct {
// 	Name string
// }
