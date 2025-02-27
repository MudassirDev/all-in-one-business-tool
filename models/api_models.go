package models

import (
	"time"

	"github.com/MudassirDev/all-in-one-business-tool/internal/database"
	"github.com/google/uuid"
)

type APICfg struct {
	DB              *database.Queries
	JWTSecretKey    string
	JWTExpiringTime time.Duration
	AuthCookieName  string
}

type UserStruct struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
