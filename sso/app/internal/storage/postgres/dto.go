package postgres

import domain "vss/sso/internal/domain/user"

type CreateUser struct {
	Login          string `db:"login"`
	HashedPassword string `db:"Password"`
}

type User struct {
	ID    string `db:"id"`
	Login string `db:"login"`
}

type UpdateUser struct {
	ID             string `db:"id"`
	Login          string `db:"login"`
	HashedPassword string `db:"Password"`
}

func CreateUserDTO(in domain.RegisterUser) CreateUser {
	return CreateUser{
		Login:          in.Login,
		HashedPassword: in.Password,
	}
}
