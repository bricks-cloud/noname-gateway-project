package http

type NotFoundError struct {
	message string
}

func NewNotFoundError(msg string) *NotFoundError {
	return &NotFoundError{
		message: msg,
	}
}

func (nfe *NotFoundError) Error() string {
	return nfe.message
}
