package core

type UnauthorizedError struct {
	Message string
}

type BadRequestError struct {
	Message string
}

func (e UnauthorizedError) Error() string {
	return e.Message
}

func (e BadRequestError) Error() string {
	return e.Message
}
