package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey"`
	Username   string `gorm:"size:64"`
	Password   string `gorm:"size:255"`
	CreditCard CreditCard
}

func (u *User) TableName() string {
	return "gormbasics1.users"
}

type Note struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey"`
	Name    string `gorm:"size:255"`
	Content string `gorm:"type:text"`
	UserId  uint64 `gorm:"index"`
	User    User 
}

func (u *Note) TableName() string {
	return "gormbasics1.note"
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint64
}

func (u *CreditCard) TableName() string {
	return "gormbasics1.creditcard"
}
