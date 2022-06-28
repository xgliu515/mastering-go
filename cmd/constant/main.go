package main

import (
	"fmt"
)

const gc = 1

func test1() {
	const c = 1

	fmt.Printf("%p - %p\n", &gc, &c)
}

func main() {
	test1()
}
