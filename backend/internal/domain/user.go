package domain

import (
	"time"

	"github.com/MatheusHenrique129/application-in-go/internal/model"
)

// :: Input structs

type URIUser struct {
	UserID string `uri:"user_id" json:"-" yaml:"-" binding:"required"`
}

type CreateUser struct {
	Name        string `json:"name" yaml:"name" example:"john" binding:"required,min=2,max=70" validate:"required"`
	Email       string `json:"email" yaml:"email" validate:"required,min=2,max=50,valid_email"`
	BirthDate   string `json:"birth_date" yaml:"birth_date" validate:"required"`
	PhoneNumber string `json:"phone_number" yaml:"phone_number" validate:"required"`
	Password    string `json:"password" yaml:"password" validate:"required,min=5,max=15"`
	CPF         string `json:"cpf" yaml:"cpf" validate:"required"`
	Address     string `json:"address" yaml:"address" validate:"required"`
	Gender      string `json:"gender" yaml:"gender" validate:"required"`
}

// :: Output structs

type User struct {
	UserID      int64  `json:"user_id" yaml:"user_id"`
	Name        string `json:"name" yaml:"name"`
	Email       string `json:"email" yaml:"email"`
	BirthDate   string `json:"birth_date" yaml:"birth_date"`
	PhoneNumber string `json:"phone_number" yaml:"phone_number"`
	CPF         string `json:"cpf" yaml:"cpf"`
	Address     string `json:"address" yaml:"address"`
	Gender      string `json:"gender" yaml:"gender"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func UserToCreateUserModel(req *CreateUser) *model.User {

	return &model.User{
		Name:        req.Name,
		CPF:         req.CPF,
		Email:       req.Email,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Gender:      req.Gender,
		Password:    req.Password,
		BirthDate:   req.BirthDate,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func UserToCreateUserDomain(req *model.User) *User {

	return &User{
		UserID:      req.ID,
		Name:        req.Name,
		CPF:         req.CPF,
		Email:       req.Email,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Gender:      req.Gender,
		BirthDate:   req.BirthDate,
		CreatedAt:   req.CreatedAt.Format("02.01.2006 03:04:05"),
		UpdatedAt:   req.UpdatedAt.Format("02.01.2006 03:04:05"),
	}
}
