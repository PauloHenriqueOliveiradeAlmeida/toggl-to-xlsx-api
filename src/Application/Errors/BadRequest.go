package Errors

type BadRequest struct {
	Message string `json:"message"`
}

func (this *BadRequest) Error() string {
	return this.Message
}
