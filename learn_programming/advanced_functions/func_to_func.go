package main

import "fmt"

func addOne(a int) int {
	return a + 1
}

func addTwo(a int) int {
	return a + 2
}

// in printOperations we pass a function that takes one
// int and returns one int, therefore we can easily pass
// addOne and addTwo as both are taking only one parameter and
// returning one parameter, if there is any difference, we won't
// be able to pass the parameters
func printOperations(a int, f func(int) int) {
	fmt.Println(f(a))
}

func main() {
	printOperations(1, addOne)
	printOperations(1, addTwo)
}
