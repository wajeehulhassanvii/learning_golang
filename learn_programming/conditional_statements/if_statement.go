package main

import "fmt"

func main() {

	a := 12

	if a > 0 {
		fmt.Println(a)
	}

	if true {
		fmt.Println(a)
	}

	// we can ddeclare common variable for "if" statement
	// and use it in both if and else block like below
	// the defined variable will remain for loop only
	if b := a / 2; b > 5 {
		fmt.Println("a is bigger than 5:", a, b)
	} else {
		fmt.Println("a is smaller than 5:", a, b)
	}
}
