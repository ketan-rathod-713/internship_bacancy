package school

// const table_schema = "gormschool"

type School struct {
	Id   uint64 `gorm:"primaryKey"`
	Name string
}

type Parent struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string
	Students []Student // they can have multiple childs
}

type Student struct {
	Id       uint64 `gorm:"primaryKey"`
	Name     string
	SchoolId uint64 // foreign key SchoolId References School(Id)
	School   School
	// Parents  []Parent // there can be 2 parents
}

func (s *Student) TableName() string {
	return "gormschool.student"
}

func (s *School) TableName() string {
	return "gormschool.school"
}
