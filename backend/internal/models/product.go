package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Images      []*Photo
	Title       string
	Description string
	Price       float64
	Quantity    *uint
}

type Photo struct {
	gorm.Model
	PhotoUrl      string
	FeaturedPhoto bool
}
