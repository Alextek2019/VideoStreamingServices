package user

const (
	queryCreateUser = `
		insert into users.user(login, password) values ($1, $2);
`

	queryUpdateUser = ``

	queryGetUser = ``

	queryDeleteUser = ``
)
