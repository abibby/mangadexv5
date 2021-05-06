package mangadexv5

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type APIError struct {
	ID      string          `json:"id"`
	Status  int             `json:"status"`
	Title   string          `json:"title"`
	Detail  string          `json:"detail"`
	Context json.RawMessage `json:"context"`
}

type APIErrorList []*APIError

func (e APIErrorList) Error() string {
	return e[0].Title
}

type APIResponseError struct {
	response *http.Response
}

func NewAPIResponseError(response *http.Response) error {
	return &APIResponseError{response: response}
}

func (e *APIResponseError) Error() string {
	defer e.response.Body.Close()
	body, err := ioutil.ReadAll(e.response.Body)
	if err != nil {
		body = []byte{}
	}
	return fmt.Sprintf("request failed with status code %d: %s", e.response.StatusCode, body)
}
