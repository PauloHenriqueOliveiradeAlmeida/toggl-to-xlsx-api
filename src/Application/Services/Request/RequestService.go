package Request

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func Send[Response interface{}](client *http.Client, request *Request) (Response, error) {

	var result Response

	req, error := http.NewRequest(string(request.method), request.endpoint, bytes.NewBuffer(request.serializeBody()))
	if error != nil {
		return result, error
	}

	req.Header.Add("Content-Type", "application/json")
	for key, value := range request.headers {
		req.Header.Add(key, value)
	}
	response, error := client.Do(req)
	if error != nil {
		return result, error
	}

	if response.StatusCode >= 400 {
		return result, errors.New(http.StatusText(response.StatusCode))
	}

	defer response.Body.Close()
	responseBody, error := io.ReadAll(response.Body)
	if error != nil {
		return result, error
	}

	error = json.Unmarshal(responseBody, &result)

	if error != nil {
		return result, error
	}

	return result, nil
}
