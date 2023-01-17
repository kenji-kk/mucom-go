package models 

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id        uuid.UUID		 
	UserName  string `binding:"required,min=5,max=30"`
	Password  string `binding:"required,min=6,max=30"`
	Email     string `binding:"required,min=5,max=100"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
	Salt      []byte `json:"-"`
	HashedPassword []byte `json:"-"`
}
