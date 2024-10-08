package user

import (
	"github.com/google/uuid"
	"regexp"
)

type RegisterUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (r *RegisterUser) Validate() bool {
	loginRegexp := regexp.MustCompile(`^[a-zA-Z0-9]{5,10}$`)
	passwordRegexp := regexp.MustCompile(`^[a-zA-Z0-9]{5,10}$`)

	return loginRegexp.MatchString(r.Login) &&
		passwordRegexp.MatchString(r.Password)
}

type User struct {
	ID uuid.UUID
}
