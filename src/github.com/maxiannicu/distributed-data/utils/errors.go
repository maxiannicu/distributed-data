package utils

type invalidArgumentError struct {
	message string
}

func NewInvalidArgumentError(message string) error {
	return invalidArgumentError{
		message:message,
	}
}

func (err invalidArgumentError) Error() string {
	return err.message
}