package errors


type Error interface {
	NotFound() map[string]string
	BadRequest() map[string]string
	Unauthorized() map[string]string
}

type ErrorHandler struct {

}

func (err ErrorHandler) NotFound() map[string]string {
	return map[string]string{
		"message": "Object is not found",
	}
}

func (err ErrorHandler) BadRequest() map[string]string{
	return map[string]string{
		"message": "Bad request",
	}
}

func (err ErrorHandler) Unauthorized() map[string]string{
	return map[string]string{
		"message": "Unauthorized",
	}
}



