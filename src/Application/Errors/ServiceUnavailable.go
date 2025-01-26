package Errors

type ServiceUnavailable struct {
	Message string `json:"message"`
}

func (this *ServiceUnavailable) Error() string {
	return this.Message
}
