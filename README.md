# HTTP From Input

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