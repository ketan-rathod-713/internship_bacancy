package models

type Student struct {
	Id        uint64    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	City      string    `json:"city"`
	Pincode   string    `json:"pincode"`
	BirthDate string    `json:"birth_date"`
	SportId   uint64    `json:"sport_id"`                                                   // it will create foreign key here references to Id of Sport
	Sport     Sport     `json:"sport"`                                                      // for reference belongs to relation
	Parents   []*Parent `json:"parents" gorm:"many2many:gormschoolproject.student_parent;"` // forword many to many
}

func (s *Student) TableName() string {
	return "gormschoolproject.student"
}

type Parent struct {
	Id       uint64     `gorm:"primaryKey" json:"id"`
	Name     string     `json:"name"`
	Relation string     `json:"relation"`                                                    // father, mother
	Students []*Student `json:"students" gorm:"many2many:gormschoolproject.student_parent;"` // backword join many to many
}

func (p *Parent) TableName() string {
	return "gormschoolproject.parent"
}

// add one belongs to relation

type Sport struct {
	Id   uint64 `gorm:"primaryKey"`
	Name string
}

func (s *Sport) TableName() string {
	return "gormschoolproject.sport"
}

type StudentParentRelationData struct {
	StudentId uint64 `json:"student_id"`
	ParentId  uint64 `json:"parent_id"`
}
