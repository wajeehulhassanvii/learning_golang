package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func main() {
	// assign function to local variable
	myAddOne := addOne
	fmt.Println(addOne(1))
	fmt.Println(myAddOne(1))

	// assign a function to variable on the fly
	myAddFunc2 := func(a int) int {
		return a + 2
	}

	fmt.Println(myAddFunc2(2))
}
