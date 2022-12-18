package model

type Product struct {
	Images      []*Photo
	Title       string
	Description string
	Price       float64
	Quantity    *uint
}

type Photo struct {
	PhotoUrl      string
	FeaturedPhoto bool
}
