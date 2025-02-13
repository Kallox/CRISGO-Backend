package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey autoIncrement" json:"id"`
	Username string `gorm:"not null " json:"username"`
	Fullname string `gorm:"not null" json:"fullname"`
	Password string `gorm:"not null" json:"password"`
	Email    string `gorm:"not null unique" json:"email"`
	Gender   string `gorm:"not null" json:"gender"`
}
