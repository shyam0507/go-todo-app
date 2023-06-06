package models

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/shyam0507/todo-app/pkg/config"
)

type Todo struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Task        string    `gorm:"type:text;not null" json:"task" validate:"required"`
	Description string    `gorm:"type:text" json:"description"`
}

func (t *Todo) Validate() error {
	validate := validator.New()
	if err := validate.Struct(t); err != nil {
		return err
	}
	return nil
}

func (t *Todo) Save() (*Todo, error) {
	tx := config.DB.Create(&t)
	return t, tx.Error
}
