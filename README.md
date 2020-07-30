[![Coverage Status](https://coveralls.io/repos/github/johnmackenzie91/httptestfixtures/badge.svg?branch=master)](https://coveralls.io/github/johnmackenzie91/httptestfixtures?branch=master)

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

Allowing for easier unit testing of api clients.

You can use curl to download to hit your api and download the entire response

```
curl -i https://yourapi.com/endpointone > ./testdata/endpointone.txt
```
Then use that response as a accurate representation of how your endpoint will respond and create a unit test around that.
Please see the [examples](https://github.com/johnmackenzie91/httptestfixtures/tree/master/examples).

