package main

import (
	"fmt"

	"github.com/jonbodner/leftpad"
)

func main() {
	fmt.Println(leftpad.Format("hello", 15))
	fmt.Println(leftpad.FormatRune("hello", 15, "®"))
}

// no circular imports
// each package cannot import from one another
