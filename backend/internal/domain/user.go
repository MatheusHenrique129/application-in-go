package domain

type URIUser struct {
	UserID string `uri:"user_id" example:"170" binding:"-" validate:"required"`
}

type User struct {
	UserID      int64  `json:"user_id" yaml:"user_id" example:"170"`
	Name        string `json:"name" yaml:"name" example:"john" binding:"required,min=2,max=70" validate:"required"`
	Email       string `json:"email" yaml:"email" validate:"required,min=2,max=50,valid_email"`
	BirthDate   string `json:"birth_date" yaml:"birth_date" validate:"required"`
	PhoneNumber string `json:"phone_number" yaml:"phone_number" validate:"required"`
	Password    string `json:"password" yaml:"password" validate:"required,min=8,max=15"`
	CPF         string `json:"cpf" yaml:"cpf" validate:"required"`
	Address     string `json:"address" yaml:"address" validate:"required"`
	Gender      string `json:"gender" yaml:"gender" validate:"required"`
}
