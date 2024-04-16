package models

type User struct {
	Id     int    `gorm:"primaryKey,autoIncrement"`
	Name   string `gorm:"not null"`
	Gender string
	Age    string `gorm:"check:age>10"`
}
