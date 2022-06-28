package main

import (
	"fmt"
	"runtime"
)

type data struct {
	i, j int
}

const N = 40000000

func test1() {
	var s []data

	for i := 1; i < N; i++ {
		v := int(i)
		s = append(s, data{v, v})
	}

	runtime.GC()
	_ = s[0]
}

func test2() {
	s := make(map[int]*int)

	for i := 1; i < N; i++ {
		v := int(i)
		s[v] = &v
	}

	runtime.GC()
	_ = s[0]
}

func main() {
	test1()
	fmt.Println("----------------------")
	test2()
}
