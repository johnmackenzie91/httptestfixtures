package httpfrominput

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseStatusCode(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expectedOut statusCode
		expectedErr error
	}{
		{
			name:  "http version 1.somthing ok",
			input: "HTTP/1.1 200 OK",
			expectedOut: statusCode{
				StatusCode: 200,
				Status:     "200 OK",
			},
		},
		{
			name:  "http version 2",
			input: "HTTP/2 200 ",
			expectedOut: statusCode{
				StatusCode: 200,
				Status:     "200 OK",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// arrange
			// act
			res, err := parseStatusCode(tc.input)

			// assert
			assert.Equal(t, tc.expectedOut, res)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
