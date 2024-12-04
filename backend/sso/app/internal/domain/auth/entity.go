package auth

import (
	"time"

	"github.com/gofrs/uuid"
)

type SignInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignInResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	AccessToken uuid.UUID `json:"access_token"`
	TTL         time.Time `json:"ttl"`
}

type SignOutRequest struct {
	UserID      uuid.UUID `json:"user_id"`
	AccessToken string    `json:"access_token"`
}
