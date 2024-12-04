package errors

var (
	// User errors
	ErrBodyParsing  = NewCustomError("Unable to parse request body", 400, true, "ERR_BODY_PARSING")
	ErrUserID       = NewCustomError("Invalid user ID", 400, true, "ERR_USER_ID")
	ErrUserNotFound = NewCustomError("Could not find user", 400, true, "ERR_USER_NOT_FOUND")

	// Auth errors
	ErrInvalidAuth = NewCustomError("Invalid auth", 400, true, "ERR_INVALID_AUTH")
)
