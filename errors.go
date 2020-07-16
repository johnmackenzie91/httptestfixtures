package httptestfixtures

import "errors"

// ErrUnableToParseStatusCode fails if unable to parse fire line of file
var ErrUnableToParseStatusCode = errors.New("unable to parse status code")
var ErrUnableToParseHeader = errors.New("unable to parse header")
