package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/geekysaurabh001/gomongo/routers"
)

func main() {
	fmt.Println("Welcome to Golang MongoDB API")
	fmt.Println("Server is getting started...")
	r := routers.Router()
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at port 4000: http://localhost:4000")
}
