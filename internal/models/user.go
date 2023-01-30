package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id             uuid.UUID `json:"id" db:"id" validate:"omitempty"`
	UserName       string    `json:"user_name" db:"user_name" binding:"required,min=5,max=30"`
	Password       string    `json:"password" db:"password" binding:"required,min=5,max=30"`
	Email          string    `json:"email" db:"email" binding:"required,min=5,max=100"`
	CreatedAt      time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at,omitempty" db:"updated_at"`
	Salt           []byte    `json:"-" db:"salt"`
	HashedPassword []byte    `json:"-" db:"hashed_password"`
}
