package models

type DocumentHistory struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Timestamp     string `gorm:"type:timestamp"`
	Path          string `gorm:"size:255"`
	VersionNumber int
	UserID        uint     `gorm:"column:user_id"`
	User          UserAuth `gorm:"foreignKey:UserID;references:ID"`
	DocumentID    uint     `gorm:"column:document_id"`
	Document      Document `gorm:"foreignKey:DocumentID;references:ID"`
}
