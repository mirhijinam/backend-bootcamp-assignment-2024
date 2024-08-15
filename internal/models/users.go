package models

type Role string

const (
	RoleClient    Role = "client"
	RoleModerator Role = "moderator"
)

type User struct {
	ID    string
	Email string
	Role  Role
}

type UserClaims struct {
	ID   string
	Role Role
}
