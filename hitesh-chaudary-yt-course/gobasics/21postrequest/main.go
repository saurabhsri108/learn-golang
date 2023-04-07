package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url string = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Welcome to post request")
	PerformPostJsonRequest()
}

func PerformPostJsonRequest() {
	requestBody := strings.NewReader(`
		{
			"title": "Post 101",
			"body": "Lorem up your @ss",
			"userId": 1
		}
	`)

	response, err := http.Post(url, "application/json, charset=UTF-8", requestBody)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
