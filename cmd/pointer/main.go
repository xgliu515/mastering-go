package main

import "fmt"

func main() {
	i := 1
	s := []int{1, 2, 3}

	fmt.Printf("%p, %p\n", s, &i)
	fmt.Printf("%T, %T\n", s, &i)
}
