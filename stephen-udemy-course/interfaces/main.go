package main

/**
type _interfaceName_ interface {
	functionName(argType1, argType2) (returnType1, returnType2)
}

concrete types: map, struct, int, string, englishBot
interface types: bot

1. Interfaces are not generic types. Golang doesn't have generic types.

2. Interfaces are 'implicit'. Meaning, we don't have to tell eb and sb that it is linked to bot. Golang takes care of it for us. Good and bad both.

3. Interfaces are a contract to help us manage types. It doesn't check anything or provide any verification. Garbage In -> Garbage Out. If our custom type's implementation of a function is broken then interfaces won't help us!

4. Interfaces are tough. Step #1 is understanding how to read them: Understand how to read interfaces in the standard lib. Writing your own interfaces is tough and requires experiences.
*/

// type bot interface {
// 	getGreeting() string
// }
// type englishBot struct{}
// type spanishBot struct{}

func main() {
	// eb := englishBot{}
	// sb := spanishBot{}

	// printGreeting(eb)
	// printGreeting(sb)
	// httpMain()
	// assignment1()
	assignment2()
}

// func printGreeting(b bot) {
// 	fmt.Println(b.getGreeting())
// }

// func (englishBot) getGreeting() string {
// 	return "Hi There!"
// }

// func (spanishBot) getGreeting() string {
// 	return "Hola"
// }
