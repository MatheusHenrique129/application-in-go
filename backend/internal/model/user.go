package model

import "time"

type User struct {
	ID          int64     `db:"id"`
	Name        string    `db:"full_name"`
	CPF         string    `db:"cpf"`
	Email       string    `db:"email"`
	Address     string    `db:"address"`
	PhoneNumber string    `db:"phone_number"`
	Gender      string    `db:"gender"`
	Password    string    `db:"password"`
	BirthDate   string    `db:"birth_date"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
