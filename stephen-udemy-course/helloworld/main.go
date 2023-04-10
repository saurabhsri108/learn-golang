// There are 2 types of packages in Go: executables and libraries (re-usables). main package name is reserved for executables. This is the package that will be compiled and run.
package main

// Importing fmt package to print to the console
import "fmt"

// main function is the entry point for the program. This is where the program starts executing.
func main() {
	fmt.Println("Hello, world.")
}
