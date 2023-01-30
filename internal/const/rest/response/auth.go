package response

import (
	"github.com/kenji-kk/mucom-go/internal/models"
)

type SignupResponse struct {
	User *models.User `json:"user"`
	JWS  string       `json:"jws"`
}
