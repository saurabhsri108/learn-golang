package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=asdasradf"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println("Scheme:", result.Scheme)       // https
	fmt.Println("User:", result.User)           // __blank__
	fmt.Println("Host:", result.Host)           // lco.dev:3000
	fmt.Println("Path:", result.Path)           // /learn
	fmt.Println("Port:", result.Port())         // 3000
	fmt.Println("Rawquery:", result.RawQuery)   // coursename=reactjs&paymentid=asdasradf
	fmt.Println("Rawpath:", result.RawPath)     // __blank__
	fmt.Println("Hostname:", result.Hostname()) // lco.dev

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams) // url.Values
	fmt.Println(qparams["coursename"])                        // reactjs

	for _, val := range qparams {
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
