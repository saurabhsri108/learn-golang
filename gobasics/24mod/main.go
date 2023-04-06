package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to go mod")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func greeter() {
	fmt.Println("Hey there! mod users")
}

func serveHome(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}

/**
go mod init github.com/username/module_name
go mod tidy (indirect -> direct in mod file)
go mod verify (to verify sum hashes to protect from attacks)
go list -m all (list all the packages in mod file)
go list -m -versions github.com/username/mod_name (list all the versions of the mod_name)
go mod why github.com/username/mod_name (list the modules dependent on this mod_name)
go mod graph (expensive command but shows the dependency graphs)
go mod edit -go go_version (to edit the mod file and update go version)
go mod vendor
go run -mod=vendor main.go (uses the package in from local folder created after the go mod vendor command)
*/
