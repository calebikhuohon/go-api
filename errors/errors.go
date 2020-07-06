package errors

//Type describes the kind of error and roughly translates to a HTTP status code
//for client errors
type Type string

const (
	//TypeBadRequest is used for HTTP 400-like errors
	TypeBadRequest Type = "bad_request_error"

	//TypeNotFound is used for HTTP 404
	TypeNotFound Type = "not_found_error"
)

//AppError is an implementation of error with types to
//differentiate client and server errors
type AppError struct {
	text string
	errType Type
}

func (e AppError) Error() string  {
	return  e.text
}

func (e AppError) Type() Type  {
	return e.errType
}

func Error(text string) error  {
	return &AppError{
		text: text,
		errType: TypeBadRequest,
	}
}

func NotFound(entity string) error {
	return &AppError{
		text: entity + "not found",
		errType: TypeNotFound,
	}
}
