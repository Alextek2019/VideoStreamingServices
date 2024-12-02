package user

const (
	queryCreateUser = `
		INSERT INTO users.user(login, password) VALUES ($1, $2)
		RETURNING id    AS id,
		  login 		AS login;
`

	queryUpdateUser = ``
)
