package schemas

import "github.com/google/uuid"

type User struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Role     string    `json:"role"` // TODO: should be ENUM
}
