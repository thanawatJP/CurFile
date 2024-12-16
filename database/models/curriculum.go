package models

type Curriculum struct {
	ID       uint      `gorm:"primaryKey;autoIncrement"`
	Name     string    `gorm:"size:255"`
	Subjects []Subject `gorm:"many2many:curriculum_subject"`
}
