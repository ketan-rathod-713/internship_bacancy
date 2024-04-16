package main

import (
	"fmt"
	"gorm/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	connString := "host=localhost user=bacancy password=admin dbname=bacancy port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // For Logging Purposes
	})
	if err != nil {
		log.Println("ERROR connecting database")
	}

	// SCHEMA INITIALIZATION
	err = db.Exec("CREATE SCHEMA IF NOT EXISTS gormbasics").Error
	if err != nil {
		panic("failed to create schema")
	}

	// Migrating Schema if not created
	err = db.AutoMigrate(&models.Student{})
	if err != nil {
		log.Fatal(err)
	}

	// Inserting Data
	result := db.Create(&models.Student{Name: "Ketan", Class: 12})
	log.Println(result.RowsAffected)
	log.Println(result.Error)

	// We can also create multiple records with create by passing slice of struct
	// we can also pass pointer to data in create.

	// CREATE record with only selected fields by using Select // TODO: why only address works here.
	db.Select("name").Create(&models.Student{Name: "Ketan", Class: 10})

	// BATCH INSERT // EFICIENCY
	// To efficiently insert large number of records, pass a slice to the Create method. GORM will generate a single SQL statement to insert all the data and backfill primary key values, hook methods will be invoked too. It will begin a transaction when records can be split into multiple batches.

	db = db.Session(&gorm.Session{CreateBatchSize: 2}) // means if 100 records then it will divide it in 50 batches and do 50 insertions at same time.

	// SEE LOGS IT WILl PRINT 3 TIMES TO INSERT THIS DATA
	var students []models.Student = []models.Student{
		{Name: "Aman", Class: 9},
		{Name: "Rahul", Class: 30},
		{Name: "Rahul Sharma"},
		{Name: "Tridip Chavda", Class: 30},
		{Name: "Tridip Chavda 2", Class: 20},
	}

	// if anyone fails then whole batch will not get inserted ha ha // its about transactions :)
	db.Create(students)

	log.Println("Id Generated So Far Is As : ")
	for _, val := range students {
		fmt.Println(val.ID)
	}

	// Lets try with one record only

	db.Create(&models.Student{Name: "AMAN created", Class: 20})
	db.Create(&models.Student{Name: "KETAN not created", Class: 26})

	// Create From Map
	// GORM supports create from map[string]interface{} and []map[string]interface{}{}
	// It doesnt create default values. :)) TODO:
	db.Model(&models.Student{}).Create(map[string]interface{}{
		"Name":  "jinhu",
		"Class": 12,
	})

	// TODO: Create From SQL Expression/Context Valuer

	// CREATE WITH ASSOCIATIONS TODO:

	// Default values with gorm:"default: 122"

	// NOTE Any zero value like 0, '', false won’t be saved into the database for those fields defined default value, you might want to use pointer type or Scanner/Valuer to avoid this, for example:

	//TODO:
	// type User struct {
	// 	gorm.Model
	// 	Name string
	// 	Age  *int           `gorm:"default:18"`
	// 	Active sql.NullBool `gorm:"default:true"`
	//   }

	a := 0 // AGAR AGE POINTER TYPE KA NA HOTA TO YE DATA SAVE NAHI HOTA HA HA  :)
	// another option is to set data in hooks after checking isnerted data
	db.Create(&models.Student{Age: &a})
	db.Create(&models.Student{})

	// NOTE You have to setup the default tag for fields having default or virtual/generated value in database, if you want to skip a default value definition when migrating, you could use default:(-), for example: TODO:

	// Upsert / On Conflict TODO: TODO:

}

/*
GORM uses the field with the name ID as the table’s primary key by default.
You can set other fields as primary key with tag primaryKey
Set multiple fields as primary key creates composite primary key
integer PrioritizedPrimaryField enables AutoIncrement by default, to disable it, you need to turn off autoIncrement for the int fields:

type Product struct {
  CategoryID uint64 `gorm:"primaryKey;autoIncrement:false"`
  TypeID     uint64 `gorm:"primaryKey;autoIncrement:false"`
}

Pluralized Table Name by default

TO change it = implement Tabler interface
type Tabler interface {
  TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
  return "profiles"
}

Above doesnt allow dynamic names, it will cashed in further operations.

for using dynamic names use scope instead.

gorm Provides predefined struct
// gorm.Model definition
type Model struct {
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

Also we can have field level permissions // TODO:

For gorm.Model it will embedd directly but for normal struct we have to define it using tags
type Blog struct {
  ID      int
  Author  Author `gorm:"embedded"` // Auther struct is embedded in it
  Upvotes int32
}

GORM TAGS
- case insensitive
- camelCase is prefered

GORM can also be connected with existing database connection // for eg. from sql one // TODO:

We are using pgx as postgres’s database/sql driver, it enables prepared statement cache by default, We can disable it TODO:

GORM using database/sql to maintain connection pool


*/
