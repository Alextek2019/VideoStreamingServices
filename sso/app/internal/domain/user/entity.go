package user

import (
	"github.com/google/uuid"
	"regexp"
)

type RegisterUser struct {
	login    string
	password string
}

func (r *RegisterUser) Validate() bool {
	loginRegexp := regexp.MustCompile(`^[a-zA-Z0-9]{5,10}$`)
	passwordRegexp := regexp.MustCompile(`^[a-zA-Z0-9]{5,10}$`)

	return loginRegexp.MatchString(r.login) &&
		passwordRegexp.MatchString(r.password)
}

type User struct {
	ID uuid.UUID
}
