# HTTP Test Fixtures

The library converts a http response file;

```
HTTP/1.1 200 OK
Date: Sun, 10 Oct 2010 23:26:07 GMT

{"msg": "Hello world!"}
```

Into a *http.Response struct;

```
&http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Header: http.Header{
			"Date": []string{"Sun, 10 Oct 2010 23:26:07 GMT"},
		},
		Body: ioutil.NopCloser(strings.NewReader(`{"msg": "Hello world!"}`)),
	}
```

Allowing for easier unit testing, please see [examples](https://github.com/johnmackenzie91/httptestfixtures/tree/master/examples).