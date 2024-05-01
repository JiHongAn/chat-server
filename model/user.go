package model

type User struct {
	Name     string `gorm:"type:varchar(100);id"`
	Email    string `gorm:"type:varchar(100);uniqueIndex"`
	Password string `gorm:"type:varchar(255)"`
}
