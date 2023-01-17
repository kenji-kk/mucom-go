package utils 

import (
	"github.com/go-playground/validator/v10"
)

var validate *validator.validate

func init() {
	validate = validator.New()
}
