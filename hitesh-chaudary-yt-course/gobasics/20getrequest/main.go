package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const url string = "https://jsonplaceholder.typicode.com/posts/1"

func main() {
	fmt.Println("Welcome to get request")
	PerformGetRequest()
}

func PerformGetRequest() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)
	content, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println(responseString.String())
	fmt.Println(byteCount)
}
