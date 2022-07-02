package main

import (
	"fmt"
	"time"
)

func test0() {
	fmt.Println("Hello", "world!")
}

func test1() {
	a:=1
	fmt.Printf("%v\n", &a)
	go func(a *int) {
		for i := 0; i < 3; i++ {
			fmt.Printf("%v - %d\n", a, *a)
			time.Sleep(1 * time.Second)
		}
	} (&a)

	time.Sleep(1 * time.Second)
	a = 2
	time.Sleep(3 * time.Second)
}

func test2() {
	a:=1
	fmt.Printf("%v\n", &a)
	go func(a int) {
		for i := 0; i < 3; i++ {
			fmt.Printf("%v - %d\n", &a, a)
			time.Sleep(1 * time.Second)
		}
	} (a)

	time.Sleep(1 * time.Second)
	a = 2
	time.Sleep(3 * time.Second)
}


func main() {
	test1()
	test2()
}