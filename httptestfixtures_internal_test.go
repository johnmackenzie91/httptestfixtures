package httptestfixtures

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseStatusCode(t *testing.T) {
	testCases := []struct {
		input       string
		expectedOut statusCode
		expectedErr error
	}{
		{
			input: "HTTP/1.1 200 OK",
			expectedOut: statusCode{
				StatusCode: 200,
				Status:     "200 OK",
			},
		},
		{
			input: "HTTP/2 200 ",
			expectedOut: statusCode{
				StatusCode: 200,
				Status:     "200 OK",
			},
		},
		{
			input: "HTTP/2 404",
			expectedOut: statusCode{
				StatusCode: 404,
				Status:     "404 Not Found",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			// arrange
			// act
			res, err := parseStatusCode(tc.input)

			// assert
			assert.Equal(t, tc.expectedOut, res)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
