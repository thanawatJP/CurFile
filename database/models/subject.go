package models

type Subject struct {
	ID          uint         `gorm:"primaryKey;autoIncrement"`
	Code        string       `gorm:"size:100;unique"`
	Name        string       `gorm:"size:100"`
	Curriculums []Curriculum `gorm:"many2many:curriculum_subject"`
}
