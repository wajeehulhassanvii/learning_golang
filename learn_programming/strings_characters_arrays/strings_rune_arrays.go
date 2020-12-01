package main

import "fmt"

func main() {
	var s string
	fmt.Println(s)
	s = "hello world"
	fmt.Println(s)

	s1 := "shortcut"
	fmt.Println(s1)

	// `` MULTI LINE strings are enclosed with back tick

	// any unicode, including emoji
	s3 := ` string
				s3
	`
	fmt.Println(s3)

	// we can access each char of a string with its char index
	fmt.Println(s3[2])

	// SUBSTRINGS are like python
	fmt.Println(s3[0:4])

	// get length of string like python
	fmt.Println(len(s3))

	// NON-ASCII chars make sub string complicated, EMOGIs

	// rune is  for dealing with EMOJIS
	// x1 := "Hello"
	// x2 := '😎'
	// x = x1 + string(x2)
	// fmt.Println(x)

	// ARRAYS -> fixed size
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 5

	fmt.Println(vals)

	// ARRAY has problem that it has static size, its replaced by slice in Go

}
