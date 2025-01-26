package Request

import (
	"encoding/json"
	"net/http"
)

type HttpMethod string

const (
	GET     HttpMethod = http.MethodGet
	POST    HttpMethod = http.MethodPost
	PUT     HttpMethod = http.MethodPut
	DELETE  HttpMethod = http.MethodDelete
	PATCH   HttpMethod = http.MethodPatch
	HEAD    HttpMethod = http.MethodHead
	OPTIONS HttpMethod = http.MethodOptions
)

type Request struct {
	endpoint string
	method   HttpMethod
	body     interface{}
	headers  map[string]string
}

func NewRequest(endpoint string, method HttpMethod) *Request {
	return &Request{
		endpoint: endpoint,
		method:   method,
	}
}

func (this *Request) WithBody(body interface{}) *Request {
	switch this.method {
	case GET, DELETE, HEAD, OPTIONS:
		return this
	}

	this.body = body
	return this
}

func (this *Request) WithHeaders(headers map[string]string) *Request {
	this.headers = headers
	return this
}

func (this *Request) serializeBody() []byte {
	switch this.method {
	case GET, DELETE, HEAD, OPTIONS:
		return nil
	}
	payload, error := json.Marshal(this.body)
	if error != nil {
		return nil
	}

	return payload
}
