package models 

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID    `json:"id" db:"id" validate:"omitempty"`
	UserName  string       `json:"userName" db:"user_name binding:"required,min=5,max=30"`
	Password  string       `json:"password" binding:"required,min=5,max=30"`
	Email     string       `json:"email" db:"email" binding:"required,min=5,max=100"`
	CreatedAt time.Time    `json:"createdAt,omitempty" db:"created_at"`
	UpdatedAt time.Time    `json:"updatedAt,omitempty" db:"updated_at"`
	Salt      []byte       `json:"-"`
	HashedPassword []byte  `json:"-"`
}
