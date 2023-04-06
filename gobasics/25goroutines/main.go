package main

import (
	"fmt"
	"net/http"
)

func main() {
	//  go greeter("Hello")
	//  greeter("World")
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}
	for _, website := range websiteList {
		getStatusCode(website)
	}
}

// func greeter(message string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(i+1, message)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPS in endpoint")
	}
	fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
}
