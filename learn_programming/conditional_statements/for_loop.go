package main

import "fmt"

func main() {

	// normal loop
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println(i)
		}

	}

	// special Go for loop, instead of while
	// because i declared outside foor loop and still visible
	// after the for loop
	i := 0
	for i < 10 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println(i)

	//INFINITE For loop until break condition is met

	i = 0
	for {
		fmt.Println(i)
		if i == 5 {
			break
		}
		i++
	}

	// for range loop in go is like enumerate in Python
	s := "Hello world, this is Wajeeh."
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}

	fmt.Println(string(s[8]))

}
