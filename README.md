[![Coverage Status](https://coveralls.io/repos/github/johnmackenzie91/httptestfixtures/badge.svg?branch=master)](https://coveralls.io/github/johnmackenzie91/httptestfixtures?branch=master)
[![Build Status](https://travis-ci.org/johnmackenzie91/httptestfixtures.svg?branch=master)](https://travis-ci.org/johnmackenzie91/httptestfixtures)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnmackenzie91/httptestfixtures)](https://goreportcard.com/report/github.com/johnmackenzie91/httptestfixtures)

# HTTP Test Fixtures

> Generate test fixtures that are real world responses!

When writing an api client I want to make sure the test fixtures that I am developing against are as close to real life as possible.
So why not use curl to capture real world responses and use them as test fixtures?


1. Generate the test fixtures.
```shell
$  curl https://cat-api.com/cats/1 \
     --include \
     --output "${PATH}/cats_id_exists"
     
HTTP/2 200 
server: nginx/1.18.0 (Ubuntu)
date: Mon, 27 Sep 2021 15:49:44 GMT
content-type: application/json; charset=UTF-8
vary: Accept-Encoding
access-control-allow-origin: *
cache-control: public, max-age=3600

{"data":{"id":1,"name":"Squeek","age":12}}
```

2. Load the file into an *http.Response and set this to return from the mock.
```shell
func TestClient_GetCat_ID_Exists(t *testing.T) {
    res := httptestfixtures.MustLoadRequest(t, "./testdata/cats_id_exists")
    mockClient.On("Do", mock.Anything).Return(res, nil)
  
    // initialise the client
    sut := API{
      client: mockClient
    }
    
    // run the function under test
    out, err := sut.GetCat(1)
    
    // assert
    //make your asertions
```

## Social Media

* [Twitter](https://twitter.com/JohnnMackk)
* [LinkedIn](https://www.linkedin.com/in/john-mackenzie-web-developer/)