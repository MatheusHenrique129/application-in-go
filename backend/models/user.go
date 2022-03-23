package models

import "gorm.io/gorm"

type User struct {
	//Already defines the ID, CreatedAt, UpdatedAt e DeletedAt for default
	gorm.Model
	DisplayName string
	PhotoUrl    string
	Email       string
	Password    string
}
