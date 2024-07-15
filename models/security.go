package models

type Password struct {
	ID       uint   `gorm:"primarykey"`
	Name     string `gorm:"default:null"`
	Username string `gorm:"default:null"`
	Password string `gorm:"default:null"`
	Link     string `gorm:"default:null"`
}
