package errors

var (
	// ErrBodyParsing means that invalid user input in request body
	ErrBodyParsing  = NewCustomError("Unable to parse request body", 400, true, "ERR_BODY_PARSING")
	ErrUserID       = NewCustomError("Invalid user ID", 400, true, "ERR_USER_ID")
	ErrUserNotFound = NewCustomError("Could not find user", 400, true, "ERR_USER_NOT_FOUND")
)
