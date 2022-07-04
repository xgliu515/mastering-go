package main

import "fmt"

type s1 struct {
	i int
	s string
}

type s2 struct {
	s string
	i int
}

func test1() {
	i := new(int)

	fmt.Printf("%T\n", i)
}

func main() {
	test1()
}
