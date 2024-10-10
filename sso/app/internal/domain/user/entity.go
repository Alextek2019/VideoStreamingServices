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
	loginRegexp := regexp.MustCompile(`^[a-zA-Z0-9]{5,10}$`)
	passwordRegexp := regexp.MustCompile(`^[a-zA-Z0-9]{5,10}$`)

	return loginRegexp.MatchString(r.Login) &&
		passwordRegexp.MatchString(r.Password)
}

type User struct {
	UserID uuid.UUID `json:"userID"`
	Login  string    `json:"login"`
}

type UpdateUserArgs struct {
	UserID   uuid.UUID `json:"userID"`
	Login    *string   `json:"login"`
	Password *string   `json:"password"`
}
