package akamai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// APIError exposes an Akamai OPEN Edge Grid Error.
type APIError struct {
	error
	Type        string         `json:"type"`
	Title       string         `json:"title"`
	Status      int            `json:"status"`
	Detail      string         `json:"detail"`
	Instance    string         `json:"instance"`
	Method      string         `json:"method"`
	ServerIP    string         `json:"serverIp"`
	ClientIP    string         `json:"clientIp"`
	RequestID   string         `json:"requestId"`
	RequestTime string         `json:"requestTime"`
	Warnings 	[]Warning 	   `json:"warnings"`
	Response    *http.Response `json:"-"`
	RawBody     string         `json:"-"`
}

// Warning represents less severe issue.
type Warning struct {
	Type string `json:"type"`
	Detail string `json:"detail"`
	MessageID string `json:"messageId"`
}

// Error is custom implementation of error interface.
func (error APIError) Error() string {
	return strings.TrimSpace(fmt.Sprintf("API Error: %d %s %s More Info %s",
		error.Status,
		error.Title,
		error.Detail,
		error.Type))
}

// NewAPIError creates a new API error based on a Response,
// or http.Response-like.
func NewAPIError(response *http.Response) *APIError {
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))
	return NewAPIErrorFromBody(response, body)
}

// NewAPIErrorFromBody creates a new API error, allowing you to pass in a body
//
// This function is intended to be used after the body has already been read for
// other purposes.
func NewAPIErrorFromBody(response *http.Response, body []byte) *APIError {
	var apiErr APIError

	if err := json.Unmarshal(body, &apiErr); err == nil {
		apiErr.Status = response.StatusCode
		apiErr.Title = response.Status
	}

	apiErr.Response = response
	apiErr.RawBody = string(body)

	return &apiErr
}

// IsInformational determines if a response was informational (1XX status)
func IsInformational(r *http.Response) bool {
	return r.StatusCode > 99 && r.StatusCode < 200
}

// IsSuccess determines if a response was successful (2XX status)
func IsSuccess(r *http.Response) bool {
	return r.StatusCode > 199 && r.StatusCode < 300
}

// IsRedirection determines if a response was a redirect (3XX status)
func IsRedirection(r *http.Response) bool {
	return r.StatusCode > 299 && r.StatusCode < 400
}

// IsClientError determines if a response was a client error (4XX status)
func IsClientError(r *http.Response) bool {
	return r.StatusCode > 399 && r.StatusCode < 500
}

// IsServerError determines if a response was a server error (5XX status)
func IsServerError(r *http.Response) bool {
	return r.StatusCode > 499 && r.StatusCode < 600
}

// IsError determines if the response was a client or server error (4XX or 5XX status)
func IsError(r *http.Response) bool {
	return r.StatusCode > 399 && r.StatusCode < 600
}
