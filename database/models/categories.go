package models

type Category struct {
	ID        uint       `gorm:"primaryKey;autoIncrement"`
	Name      string     `gorm:"size:100;unique"`
	Documents []Document `gorm:"many2many:document_categories"`
}
