package Mocks

import (
	"errors"
	"net/http"
)

type HttpTransportMock struct {
	Response *http.Response
	Error    error
}

func (this *HttpTransportMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return this.Response, this.Error
}

type CorrompedResponse struct {
	Message string
}

func (this *CorrompedResponse) Read([]byte) (int, error) {
	return 0, errors.New(this.Message)
}
