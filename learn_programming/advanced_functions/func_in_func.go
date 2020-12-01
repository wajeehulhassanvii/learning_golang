package main

import "fmt"

func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

// takes a function and returns a function
func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}

func addOneInThis(a int) int {
	return a + 1
}

func main() {

	b := 2
	c := 3

	// give the function that returns the function
	doubleAddOne := makeDoubler(addOneInThis)
	// pass argument to the returned function
	fmt.Println(doubleAddOne(2))

	// b is refered in myAdd1, we can access variable defined in main
	// inside the func, but we can't access the varaialbe inside the func
	// from main

	myAdd1X := func(a int) int {
		// value of the varaible outside the function will change
		// anonymous functions that does this are called closure
		c = c + 1
		return a + b
	}

	fmt.Println(myAdd1X(1))
	fmt.Println(c)

	add1 := makeAdder(1)
	add2 := makeAdder(2)

	fmt.Println(add1(1))
	fmt.Println(add2(2))
}
