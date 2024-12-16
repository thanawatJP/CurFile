package models

type Document struct {
	ID         uint       `gorm:"primaryKey;autoIncrement"`
	Name       string     `gorm:"type:text"`
	Detail     string     `gorm:"type:text"`
	Year       string     `gorm:"type:timestamp"`
	UserID     uint       `gorm:"column:user_id"`
	User       UserAuth   `gorm:"foreignKey:UserID;references:ID"`
	FileID     uint       `gorm:"column:file_id"`
	File       File       `gorm:"foreignKey:FileID;references:ID"`
	SubjectID  uint       `gorm:"column:subject_id"`
	Subject    Subject    `gorm:"foreignKey:SubjectID;references:ID"`
	Categories []Category `gorm:"many2many:document_categories"`
}
