package models

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/shyam0507/todo-app/pkg/config"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"type:varchar;not null" json:"name" validate:"required"`
	Email     string    `gorm:"type:varchar;not null; unique" json:"email" validate:"required,email"`
	Password  string    `gorm:"type:varchar;not null" json:"-" validate:"required,min=8,max=32"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func (u *User) Validate() error {
	validate := validator.New()
	if err := validate.Struct(u); err != nil {
		return err
	}
	return nil
}

func (r *LoginRequest) ValidateLoginRequest() error {
	validate := validator.New()
	if err := validate.Struct(r); err != nil {
		return err
	}
	return nil
}

func (u *User) Save() error {
	if err := u.Validate(); err != nil {
		return err
	}
	if t := config.DB.Create(&u); t.Error != nil {
		return t.Error
	}
	return nil
}

func (r *LoginRequest) Login() (*User, error) {
	if err := r.ValidateLoginRequest(); err != nil {
		return nil, err
	}
	var u User
	if t := config.DB.Find(&u, &r); t.Error != nil {
		return nil, t.Error
	}
	return &u, nil
}
