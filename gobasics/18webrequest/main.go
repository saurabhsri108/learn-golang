package main

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Webrequest in golang")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close() // must be done by developers otherwise connection remains open

	databytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
