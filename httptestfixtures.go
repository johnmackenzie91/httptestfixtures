package httptestfixtures

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

// MustLoadRequest is a helper method that loads a http response from a http response file,
// if the file does not exist func fails quickly.
func MustLoadRequest(t *testing.T, filePath string) *http.Response {
	f, err := os.Open(filePath)
	if err != nil {
		t.Fatal(err)
	}
	res, err := ResponseFromReader(f)
	if err != nil {
		t.Fatal(err)
	}
	return res
}

// ResponseFromReader a reader and attempts to parse contents into a *http.Response
func ResponseFromReader(r io.Reader) (*http.Response, error) {
	s := bufio.NewScanner(r)
	res := http.Response{
		Header: http.Header{},
	}
	for i := 0; s.Scan(); i++ {
		// are we on the first line? eg. HTTP/1.1 200 OK
		if i == 0 {
			out, err := parseStatusCode(s.Text())
			if err != nil {
				return nil, err
			}
			res.StatusCode = out.StatusCode
			res.Status = out.Status
			continue
		}

		// Have we hit a blank line? If so parse body
		// and then leave
		if s.Text() == "" {
			res.Body = parseBody(s)
			continue
		}

		// else parse headers
		h, err := parseHeader(s.Text())

		if err != nil {
			return nil, err
		}
		res.Header.Add(h.key, h.value)
	}

	return &res, nil
}

// parseBody reads the remainder of the file and concatenates them together
func parseBody(s *bufio.Scanner) io.ReadCloser {
	var sb strings.Builder
	for s.Scan() {
		sb.WriteString(s.Text())
	}
	stringReader := strings.NewReader(sb.String())
	return ioutil.NopCloser(stringReader)
}

type header struct {
	key   string
	value string
}

var rxHeader = regexp.MustCompile(`([a-zA-Z0-9\-]+): (.+)`)

func parseHeader(line string) (header, error) {
	b := rxHeader.FindAllStringSubmatch(line, -1)
	if len(b) != 1 || len(b[0]) != 3 {
		return header{}, ErrUnableToParseHeader
	}
	return header{
		key:   b[0][1],
		value: b[0][2],
	}, nil
}

type statusCode struct {
	StatusCode int
	Status     string
}

var rxStatusCode = regexp.MustCompile(`(HTTPS?)\/(\d\.?\d?)\ (\d+)| ([A-Z]+)?`)

func parseStatusCode(line string) (sc statusCode, err error) {
	b := rxStatusCode.FindAllStringSubmatch(line, -1)

	if len(b) >= 1 && len(b[0]) < 3 {
		return sc, ErrUnableToParseStatusCode
	}

	out := statusCode{
		Status: b[0][3],
	}

	if out.StatusCode, err = strconv.Atoi(b[0][3]); err != nil {
		return sc, err
	}

	out.Status = strconv.Itoa(out.StatusCode) + " " + http.StatusText(out.StatusCode)

	return out, nil
}
