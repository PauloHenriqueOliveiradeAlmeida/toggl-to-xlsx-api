package Errors

type InternalError struct {
	Message string `json:"message"`
}

func (this *InternalError) Error() string {
	return this.Message
}
