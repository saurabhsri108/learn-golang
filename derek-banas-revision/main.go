package main

import (
	"fmt"
)

var pl = fmt.Println
var pf = fmt.Printf

func main() {
	// How to take user input: fmt, bufio, log, os.Stdin
	// pl("What is your name?")
	// reader := bufio.NewReader(os.Stdin)
	// name, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
	// pl("Hello", name)

	// variables
	// var vName string = "Vasu"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "hello"
	// v4 := 2.4
	// v4 = 5.4

	// data types
	// pl(reflect.TypeOf(25))      // int
	// pl(reflect.TypeOf(3.14))    // float64
	// pl(reflect.TypeOf(true))    // bool
	// pl(reflect.TypeOf("Hello")) // string
	// pl(reflect.TypeOf('🥲'))     // int32

	// casting variables
	// cv1 := 1.5
	// cv2 := int(cv1)
	// pl(cv2) // 1

	// string to integer
	// cv3 := "50000000"
	// cv4, err := strconv.Atoi(cv3) // ASCII (String) to Integer
	// pl(cv4, err, reflect.TypeOf(cv4))

	// // integer to string
	// cv5 := 50000000
	// cv6 := strconv.Itoa(cv5) // Integer to ASCII (String)
	// pl(cv4, reflect.TypeOf(cv6))

	// // string to float
	// cv7 := "3.14"
	// if cv8, err := strconv.ParseFloat(cv7, 64); err == nil {
	// 	pl(cv8, reflect.TypeOf(cv8))
	// }

	// // formating
	// cv9 := fmt.Sprintf("%f", 3.14)
	// pl(cv9, reflect.TypeOf(cv9))

	// strings
	// sv1 := "A word"
	// replacer := strings.NewReplacer("A", "Another")
	// sv2 := replacer.Replace(sv1)
	// pl(sv2)
	// pl("Length :", len(sv2))
	// pl("Contains another :", strings.Contains(sv2, "Another"))
	// pl("First index match for letter o :", strings.Index(sv2, "o"))
	// pl("Replace :", strings.Replace(sv2, "o", "0", -1)) // -1 replaces everywhere.
	// sv3 := "\nSome Words\n"                             // \t \" \\
	// sv3 = strings.TrimSpace(sv3)
	// pl("Split :", strings.Split("a-b-c-d", "-"))
	// pl("Lower :", strings.ToLower(sv2))
	// pl("Upper :", strings.ToUpper(sv2))
	// pl("Has Prefix :", strings.HasPrefix("tacocat", "taco"))
	// pl("Has Suffix :", strings.HasSuffix("tacocat", "cat"))

	// runes - characters in golang. Runes are unicode that represent characters
	// rStr := "abcdefg"
	// pl("Rune Count:", utf8.RuneCountInString(rStr))
	// for i, runeVal := range rStr {
	// 	pf("%d : %#U : %c : %d\n", i, runeVal, runeVal, runeVal)
	// }
	// Rune Count: 7
	// 0 : U+0061 'a' : a : 97
	// 1 : U+0062 'b' : b : 98
	// 2 : U+0063 'c' : c : 99
	// 3 : U+0064 'd' : d : 100
	// 4 : U+0065 'e' : e : 101
	// 5 : U+0066 'f' : f : 102
	// 6 : U+0067 'g' : g : 103

	// maths
	// pl("5 + 4 =", 5+4)   // 9
	// pl("5 - 4 =", 5-4)   // 1
	// pl("5 * 4 =", 5*4)   // 20
	// pl("5 / 4 =", 5/4)   // 1
	// pl("5 / 4 =", 5/4.0) // 1.25
	// pl("5 / 4 =", 5.0/4) // 1.25
	// pl("5 % 4 =", 5%4)   // 1

	// // short hand inc
	// mInt := 1
	// mInt += 1
	// pl(mInt)

	// // random
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1 // 0 to 50
	// pl("Random:", randNum)

	// // math operations/functions
	// pl("Abs(-10) = ", math.Abs(-10))
	// pl("Pow(4, 2) = ", math.Pow(4, 2))
	// pl("Sqrt(16) = ", math.Sqrt(16))
	// pl("Cbrt(8) = ", math.Cbrt(8))
	// pl("Ceil(4.4) = ", math.Ceil(4.4))
	// pl("Floor(4.4) = ", math.Floor(4.4))
	// pl("Round(4.4) = ", math.Round(4.4))
	// pl("Log2(8) = ", math.Log2(8))
	// pl("Log10(100) = ", math.Log10(100))
	// // Get the log of e to the power of 2
	// pl("Log(7.389) = ", math.Log(7.389))
	// pl("Max(5,4) = ", math.Max(5, 4))
	// pl("Min(5,4) = ", math.Min(5, 4))
	// convert 90 degrees to radians
	// r90 := 90 * math.Pi / 180
	// d90 := r90 * (180 / math.Pi)
	// pf("%.2f radians = %.2f degrees\n", r90, d90)
	// pl("Sin(90) =", math.Sin(r90))
	// pl("Cos(90) =", math.Cos(r90))
	// pl("Tan(90) =", math.Tan(r90))

	// printf codes
	// %d : Integer
	// %c : Character
	// %f : Float
	// %t : Boolean
	// %s : String
	// %o : Base 8
	// %x : Base 16
	// %v : Guesses based on data type
	// %T : Type of supplied value
	// pf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1) // Stuff 1 A 3.140000 true 1 1
	// pf("%9f\n", 3.14)                                               // output with total width of 9
	// pf("%.2f\n", 3.141592)                                          // output to limit value to 2 decimal places
	// pf("%9.f\n", 3.141592)                                          // output with no decimal places but with 9 width

	// sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	// pl(sp1)

	// guessing game
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)

	// randNum := rand.Intn(50) + 1

	// for true {
	// 	pl("Guess a number between 0 and 50:")
	// 	pl("Random number is:", randNum)
	// 	reader := bufio.NewReader(os.Stdin)
	// 	guess, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	guess = strings.TrimSpace(guess)
	// 	iGuess, err := strconv.Atoi(guess)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if iGuess > randNum {
	// 		pl("Pick a Lower value")
	// 	} else if iGuess < randNum {
	// 		pl("Pick a Higher value")
	// 	} else {
	// 		pl("You guessed it")
	// 		break
	// 	}
	// }

	// default values for an array based on type: 0, 0.0, false, ""
	aNums := []int{1, 2, 3, 4, 5}
	for _, num := range aNums {
		pl(num)
	}

	var arr1 [5]int
	arr1[0] = 1
	pl(arr1)
}
