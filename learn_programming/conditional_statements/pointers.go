package main

import "fmt"

// gGo does garbage collection
// Go is strongly typed
// go has its own array and string type, we don't need pointer for that
func main() {
	a := 10
	b := &a
	c := a

	fmt.Println(a, b, *b, c)

	a = 20
	fmt.Println(a, b, *b, c)

	*b = 30
	fmt.Println(a, b, *b, c)

	c = 40
	fmt.Println(a, b, *b, c)

	var x *int // 0 value of pointer, and we can't read or write
	// as it does not allocate memory
	fmt.Println(x)
	// fmt.Println(*x)

	// instead of above block
	y := new(int) // new makes pointer and allocate memory
	fmt.Println(y)
	fmt.Println(*y)

	setTo100(&a)
	fmt.Println(a)

	// we will see that the value won't change
	// since we have a brand new pointer inside the function
	setTo200Fail(&a)
	fmt.Println(a)

}

// PASSING POINTERS TO FUNCTIONS
// value is mem location where value is stored
// we can write func that can change the value of var
// in the function that calls them

func setTo100(a *int) {
	*a = 100
}

// go is still a call by value language,

func setTo200Fail(a *int) {
	a = new(int)
	*a = 200
}
