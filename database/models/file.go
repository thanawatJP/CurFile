package models

type File struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Path string `gorm:"size:255"`
}
