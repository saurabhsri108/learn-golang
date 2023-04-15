package main

import (
	"io"
	"log"
	"os"
)

func assignment2() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	io.Copy(os.Stdout, file)

	// stat, err := file.Stat()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data := make([]byte, stat.Size())
	// _, err = file.Read(data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Contents of file:", string(data))
}
