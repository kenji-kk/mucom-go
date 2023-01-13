package repository

import (
)

type AuthRepository interface {
	Hello() string
}

type authRepository struct {

}

func NewAuthRepository () AuthRepository {
	return &authRepository{}
}

func (reAuth *authRepository) Hello() string{
	return "Hello World"
}
