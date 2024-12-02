package storage

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
	ID             string  `db:"id"`
	Login          *string `db:"login"`
	HashedPassword *string `db:"Password"`
}

func CreateUserDTO(in domain.RegisterUserArgs) CreateUser {
	return CreateUser{
		Login:          in.Login,
		HashedPassword: in.Password,
	}
}

func UpdateUserDTO(in domain.UpdateUserArgs) UpdateUser {
	return UpdateUser{
		ID:             in.UserID.String(),
		Login:          in.Login,
		HashedPassword: in.Password,
	}
}