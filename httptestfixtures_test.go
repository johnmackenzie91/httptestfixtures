package httptestfixtures_test

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/johnmackenzie91/httptestfixtures"
	"github.com/stretchr/testify/assert"
)

func TestResponseFromFile_ReturnsAsExpected(t *testing.T) {
	// arrange
	f, err := os.Open("./testdata/responses/hello_world.txt")
	if err != nil {
		t.Error(err)
	}
	// act
	res, err := httptestfixtures.ResponseFromReader(f)
	// assert
	assert.Nil(t, err)
	defer res.Body.Close()

	expected := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Date": []string{"Sun, 10 Oct 2010 23:26:07 GMT"},
		},
		Body: ioutil.NopCloser(strings.NewReader(`{"msg": "Hello world!"}`)),
	}
	assert.Equal(t, expected, res)
}

func TestResponseFromFile_ReturnsAsExpected_2(t *testing.T) {
	// arrange
	f, err := os.Open("./testdata/responses/arctic-monkeys.json")
	if err != nil {
		t.Error(err)
	}
	// act
	res, err := httptestfixtures.ResponseFromReader(f)
	// assert
	assert.Nil(t, err)
	defer res.Body.Close()

	content := strings.NewReader("{\"data\": []}")
	body := ioutil.NopCloser(content)

	expected := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Server":                 []string{"Apache"},
			"P3p":                    []string{"policyref=\"/w3c/p3p.xml\" CP=\"IDC DSP COR CURa ADMa OUR IND PHY ONL COM STA\""},
			"X-Host":                 []string{"blm-web-133"},
			"X-Content-Type-Options": []string{"nosniff"},
			"Content-Type":           []string{"application/json; charset=utf-8"},
			"Date":                   []string{"Wed, 17 Jun 2020 16:16:16 GMT"},
			"Set-Cookie":             []string{"dzr_uniq_id=dzr_uniq_id_fr44eb40a9e94c3825b1684e1de2337b25a84217; expires=Mon, 14-Dec-2020 16:16:15 GMT; Max-Age=15552000; path=/; domain=.deezer.com; secure"},
			"X-Org":                  []string{"FR"},
		},
		Body: body,
	}
	assert.Equal(t, expected, res)
}
