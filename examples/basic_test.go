package examples_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/johnmackenzie91/httptestfixtures"
)

// Do sends the http request, the response of this is what we will be stubbing
type Doer interface {
	Do(r *http.Request) (*http.Response, error)
}

// MockDoer implements the above do
type MockDoer struct {
	response *http.Response
}

// Do will return whatever we have told it to return
func (m MockDoer) Do(r *http.Request) (*http.Response, error) {
	return m.response, nil
}

func GetMessageFromEndpoint(client Doer) (string, error) {
	res, err := client.Do(&http.Request{})
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	exp := struct {
		Msg string `json:"msg"`
	}{}
	if err := json.NewDecoder(res.Body).Decode(&exp); err != nil {
		return "", nil
	}
	return exp.Msg, nil
}

func Test_GetMessageFromEndpoint(t *testing.T) {
	res := httptestfixtures.MustLoadRequest(t, "../testdata/responses/hello_world.txt")
	mock := MockDoer{
		response: res,
	}

	out, _ := GetMessageFromEndpoint(mock) // err omitted for brevity

	if out != "Hello world!" {
		t.Fatalf("unexpected output: %s\n", out)
	}
}
