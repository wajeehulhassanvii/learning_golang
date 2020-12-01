// variables are declared with var and need type
package main

import "fmt"

// outiside any function a variable can only be decalred with var
var k int = 10

func main() {
	var i = 10     // it also initializes and can be used inside or outside the func
	var j int = 10 // to specify the data type explicitly
	n := 15        // this only works inside a function and it infer the data type
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(n)

	// inside the function a variable can be declared without
	// var but with colon like below

	// to assign values we don't need colon, thats only for
	// declaration and initialization

	// NOTE all the declared variable without any value, go will assign
	// 0 value for numeric

}
