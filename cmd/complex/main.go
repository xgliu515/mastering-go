package main

import (
	"fmt"
)

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of c2: %T\n", c2)
}
