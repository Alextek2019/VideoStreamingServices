package user

import (
	"github.com/gofrs/uuid"
	"regexp"
)

type RegisterUserArgs struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (r *RegisterUserArgs) Validate() bool {
	loginRegexp := regexp.MustCompile(`^[a-zA-Z0-9\W_]{5,50}$`)
	passwordRegexp := regexp.MustCompile(`^[a-zA-Z0-9\W_]{8,50}$`)

	return loginRegexp.MatchString(r.Login) &&
		passwordRegexp.MatchString(r.Password)
}

type User struct {
	UserID uuid.UUID `json:"user_id"`
	Login  string    `json:"login"`
}

type UpdateUserArgs struct {
	UserID   uuid.UUID `json:"user_id"`
	Login    *string   `json:"login"`
	Password *string   `json:"password"`
}
