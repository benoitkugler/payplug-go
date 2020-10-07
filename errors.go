package payplug

import (
	"errors"
	"fmt"
)

var (
	// Configuration failure.
	ConfigurationError = errors.New("configuration error")

	// Trying to process a request despite the fact that the secret key was not set.
	// If this is raised, you should define a configuration first with SetSecretKey
	SecretKeyNotSet = errors.New("payplug secret key is missing")
)

// Raised when there was an unrecoverable error during the request.
// This is not an unexpected HTTP response code.
type ClientError struct {
	err error
}

func (c ClientError) Error() string {
	return fmt.Sprintf("error during request: %s", c.err)
}

type HttpError struct {
	code int
	err  string
}

func (h HttpError) Error() string {
	return fmt.Sprintf("%s: the server gave the following response: `%s`.",
		mapHttpStatusToString(h.code), h.err)
}

func mapHttpStatusToString(code int) string {
	switch code {
	case 400:
		return "bad request"
	case 401:
		return "unauthorized; please check your secret key"
	case 403:
		return "forbidden error; you are not allowed to access this resource"
	case 404:
		return "the resource you requested could not be found"
	case 405:
		return "the requested method is not supported by this resource"
	}
	if 500 <= code && code <= 599 {
		return "unexpected server error during the request"
	}
	return "unhandled HTTP error"
}

// Raised when we expected the API to have a specific format, and we got something else.
func UnexpectedAPIResponseErr(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("API response is not valid JSON: %s", err)
}