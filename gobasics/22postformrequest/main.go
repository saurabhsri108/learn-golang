package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const myurl string = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Postform Request in golang")
	PerformPostformRequest()
}

func PerformPostformRequest() {
	data := url.Values{}
	data.Add("firstname", "vasu")
	data.Add("lastname", "choudhary")
	data.Add("email", "vasu@gmail.com")

	response, err := http.PostForm(myurl, data)
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
