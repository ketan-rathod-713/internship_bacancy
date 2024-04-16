package models

type Config struct {
	PORT             string
	DB_PORT          string
	DATABASE         string
	HOST             string
	DB_USER          string
	DB_USER_PASSWORD string
	DB_SCHEMA_NAME   string
}

type PersonInfo struct {
	FirstName string `gorm:"constraint:not null"`
	LastName  string `gorm:"constraint:not null"`
	SerName   string `gorm:"constraint:not null"`
}

// parents table
type Parent struct {
	Id         int    `gorm:"primaryKey;autoIncrement"`
	FatherName string `gorm:"type:varchar(255) not null"`
	MotherName string
	Students   []*Student `gorm:"-"`
	// StudentId  int
}

// students table
type Student struct {
	Id          int `gorm:"primaryKey;autoIncrement"`
	PersonInfo  `gorm:"embedded;embeddedPrefix:parent_"`
	Parent      *Parent `gorm:"foreignKey:ParentRefer;references:Id"` // reference to above so that i can do preload when required // by default references primary key hoti he foreign table ki
	ParentRefer int     // tell Parent struct to refer this field when doing preloading
	// ParentId   int // ? It would have been valid key for FK but we are going to use custom name here

	Cousins  []*Student `gorm:"-"` // avoid it inside database // this is just for my own case
	Hobbys  []*Hobby // hobby ke sath .. one to many
}

// has one ke liye ek foreign key chahiye thi but for many to many no required.
// hobby related informations
type Hobby struct {
	Id    int    `gorm:"primaryKey"`
	Title string `gorm:"constraint:unique;"`
}

// ? Student Parent Belongs to relationship

// ? In such cases we have to have 2 fields one FK referencing the Parent table and one field to hold the data of the Parent in case we are fetching Parent data.
// ? By default here ParentId will be considered as that field else we can name it deferently and provide a json tag to counter that.

// friends table to get students friends

// ? DOCS : To define a belongs to relationship, the foreign key must exist, the default foreign key uses the owner’s type name plus its primary field name.

// ! Important

// ! for our custum logic we have to override two things
// ! First is forein key and its struct  and second is references in referencing table.
// ! by default gorm select references as second tables primary key.
