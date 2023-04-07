package main

import "fmt"

const LoginToken string = "login-token" // variable declared with capital letter is public variable

func main() {
	var name string = "Saurabh Srivastava"
	fmt.Println(name)
	fmt.Printf("Name type::: %T \n", name)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable type::: %T\n", isLoggedIn)

	var age int8 = 28
	fmt.Println(age)
	fmt.Printf("Variable type::: %T\n", age)

	var unnamed string
	var unaged int
	fmt.Println(unnamed, unaged)

	var website = "ibcoder.com"
	fmt.Println(website)
	fmt.Printf("Variable type::: %T\n", website)

	numberOfUser := 300000.0 // this is walrus operator which can be used inside the function or method only, not outside
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable type::: %T\n", LoginToken)
}
