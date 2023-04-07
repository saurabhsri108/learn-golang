package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in file"
	writeFile(content)
	readFile("./gofile.txt")
}

func writeFile(content string) {
	file, err := os.Create("./gofile.txt")
	checkNilError(err)
	length, err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Length is: ", length)
	defer file.Close()
}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is\n", strings.TrimSpace(string(databyte)))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
