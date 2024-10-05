package postgres

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
