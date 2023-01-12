package repository

import (
	"fmt"
)

type AuthRepository interface {

}

type authRepository struct {

}

func NewAuthRepository () AuthRepository {
	return &authRepository{}
}

func (a *authRepository) createUser() {
	fmt.Println("aaa")
}
