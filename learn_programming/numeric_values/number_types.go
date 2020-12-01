package main

import "fmt"

func main() {
	var i int8 = 20
	var f int32 = 5
	fmt.Println(int32(i) + f)
}

// go prevent automatically type conversion
// we  can convert any value explicitely by type casting
// we even need type conversion in int and uint
