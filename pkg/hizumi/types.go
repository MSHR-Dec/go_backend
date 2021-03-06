package hizumi

type NotFound struct {
	Message string
}

func (e NotFound) Error() string {
	return e.Message
}

type InternalServerError struct {
	Message string
}

func (e InternalServerError) Error() string {
	return e.Message
}

type BadRequest struct {
	Message string
}

func (e BadRequest) Error() string {
	return e.Message
}
