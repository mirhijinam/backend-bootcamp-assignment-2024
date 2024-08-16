package models

import "github.com/google/uuid"

type Role string

const (
	RoleClient    Role = "client"
	RoleModerator Role = "moderator"
)

type User struct {
	ID       uuid.UUID
	Email    string
	Password string
	Role     Role
}

type UserClaims struct {
	ID   uuid.UUID
	Role Role
}
