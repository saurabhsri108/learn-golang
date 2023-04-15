package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/**

Signature of Response Structure in Golang

Response Struct {
	Status: string
	StatusCode: int
	Body: io.ReadCloser
}

io.ReadCloser interface {
	Reader
	Closer
}

io.Reader interface {
	Read([]byte) (int, error)
}

io.Closer interface {
	Closer() (error)
}


Sources of input:Returns:To Print it...

1. HTTP Request Body: []flargen:func printHTTP([]flargen)
2. Text file on hard drive: []string : func printFile([]string)
3. Image file on hard drive: jpegne : func printImage(jpegne)
4. User entering text into CLI: []byte: func printText([]byte)
5. Data from analog sensor plugged into machine: []float: func printData([]float)
*/

type customLogWrite struct{}

func httpMain() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// above is equivalent to below code
	// os.Stdout implements a Writer Interface
	// resp.Body implements a Reader Interface
	// io.Copy signature is (dst Writer, src Reader)
	// io.Copy(os.Stdout, resp.Body)
	clw := customLogWrite{}
	io.Copy(clw, resp.Body)
}

func (customLogWrite) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
