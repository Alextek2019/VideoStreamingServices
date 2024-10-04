package errors

var (
	// ErrBodyParsing means that invalid user input in request body
	ErrBodyParsing = NewCustomError("Unable to parse request body", 400, true, "ERR_BODY_PARSING")
)
