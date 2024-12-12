package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Posts    []Post `json:"posts"`
}

type Post struct {
	gorm.Model
	Title  string `json:"title" gorm:"not null"`
	Body   string `json:"body"`
	UserID uint   `json:"user_id"`
}
