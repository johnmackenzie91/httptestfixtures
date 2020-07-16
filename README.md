# HTTP Test Fixtures

This library makes stubbing *http.Response structs easier;

1. Fetch the endpoint you would like to stub.
```
curl -i https://www.example.com/hello > ./testdata/endpoint.txt
```

Save the contents to ./testdata/endpoint.txt
```
HTTP/1.1 200 OK
Date: Sun, 10 Oct 2010 23:26:07 GMT

Hello world!
```

2. Load content through MustLoadRequest helper function
```go
func Test_endpoint(t *testing.T) {
    var res *http.Request
    res = httptestfixtures.MustLoadRequest(t)
    


    
}
```
Imagine you have to write a client func that speaks to a /pets endpoint who's payload looks something like;
```json
HTTP/2 200 
server: Apache

{
  "data": [
    {
      "name": "Oscar",
      "type": "dog"
    },
    {
      "name": "Percy",
      "type": "parrot"
    },
  ]
}
```

Your client handler code has to take into account status code, response body, possibly even headers.
Allows you to easily create a s

This imaginatively named library consumes a http response and returns a *http.Response object.
This can come in useful for unit tests mocks.

```go
	// arrange
	f, err := os.Open("./testdata/responses/hello_world.txt")
	if err != nil {
		t.Error(err)
	}
	// act
	res, err := httpfrominput.ResponseFromReader(f)
	// assert
	assert.Nil(t, err)

	content := strings.NewReader("Hello world!")
	body := ioutil.NopCloser(content)

	expected := &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Date": []string{"Sun, 10 Oct 2010 23:26:07 GMT"},
		},
		Body: body,
	}
	assert.Equal(t, expected, res)
```