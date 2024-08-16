package dto

import "github.com/mirhijinam/backend-bootcamp-assignment-2024/internal/models"

type NewUser struct {
	Email    string
	Password string
	Role     models.Role
}
