package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(res, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(res, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(res, "hello!")
}

func formHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(res, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(res, "POST request successful\n")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(res, "Name: %v\n", name)
	fmt.Fprintf(res, "Address: %v\n", address)
}
