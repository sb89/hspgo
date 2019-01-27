package hsp

import (
	"encoding/json"
	"net/http"
)

// Error returned from wrapper or API. If StatusCode is set, error is from API otherwise wrapper.
type Error struct {
	Message    string
	StatusCode int
	Errors     []string
}

func (e Error) Error() string {
	return e.Message
}

type hspErrorResponse struct {
	MyJourneyErrors struct {
		Errors []string `json:"errors"`
	} `json:"my_journey_errors"`
}

func getError(err error) *Error {
	return &Error{Message: err.Error()}
}

func getHTTPError(resp *http.Response) *Error {
	err := Error{
		Message:    resp.Status,
		StatusCode: resp.StatusCode,
	}

	if resp.Header.Get("Content-Type") == "application/json;charset=UTF-8" {
		var apiResp hspErrorResponse
		if json.NewDecoder(resp.Body).Decode(&apiResp) == nil {
			err.Errors = apiResp.MyJourneyErrors.Errors
		}
	}

	return &err
}
