package schemas

import "errors"

type BasicResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

var (
	ErrNoAuthorizationHeaderFound = errors.New("no authorization header found")
)
