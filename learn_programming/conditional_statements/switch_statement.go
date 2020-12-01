package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	if word == "hello" {
		fmt.Println("Hi Yourself?")
	} else if word == "goodbye" {
		fmt.Println("So long!")
	} else {
		fmt.Println("I dont know what you said")
	}

	// just like if statement, we can initialize value in switch statement
	// as well
	switch l := len(word); word {
	case "hello":
		fmt.Println("Hi yourself!", l)
	// double case, single block
	case "farewell":
	case "goodbye":
		fmt.Println("So long!", l)
	case "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("Bro! You didn't said anything.")
	}

	// ADVANCED switch, one or more boolean cases
	c := "crackerjack"
	switch l := len(word); {
	case word == "hello":
		fmt.Println("Hi yourself!", l)
	// double case, single block
	case 1 < l && l < 10, word == c:
		fmt.Println("cracker jack hmmmm", l)
	case word == "farewell":
	case word == "goodbye":
		fmt.Println("So long!", l)
	case word == "greetings":
		fmt.Println("Salutations")
	default:
		fmt.Println("Bro! wtf.")
	}

}
