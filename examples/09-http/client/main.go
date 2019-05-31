package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := http.DefaultClient

	// Making a GET request will return two variables
	// A HTTP Response struct
	// An error type (interface)
	res, err := client.Get("https://developer.cisco.com")
	if err != nil {
		panic(err)
	}
	fmt.Println("status code:", res.Status)

	// If our HTTP request is more complicated, we can construct our own HTTP Request.
	data := []byte(`{"hello": "Cisco Live!"}`)
	req, err := http.NewRequest(http.MethodPost, "https://postman-echo.com/post", bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	// configure our headers and then execute the request
	req.Header.Set("Content-Type", "application/json")
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println("status code:", res.Status)

	// defer will execute at the end of our main program
	defer res.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("body: ", string(body))

}
