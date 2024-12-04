package auth

const (
	queryValidateUser = `
		SELECT id, login
			FROM users.user
		WHERE login = $1
			AND password = $2
`
)
