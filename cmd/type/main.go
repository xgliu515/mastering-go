package main

import (
	"fmt"
)

func test1() {
	i := 1

	fmt.Printf("%T\n", i)
	fmt.Println(float32(i) / 2)
}

func test2() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{4, 5, 6}

	fmt.Printf("%T, %T\n", a1, s1)
	copy(a1[0:], s1)
	fmt.Printf("%T, %T\n", a1, s1)
	fmt.Println(a1)
}

func test3() {
	t1 := new([3]int)
	t2 := make([]int, 3)
	fmt.Printf("%T, %d, %d\n", t1, cap(t1), len(t1))
	fmt.Printf("%T, %d, %d\n", t2, cap(t2), len(t2))
}

func test4() {
	s1 := []int{1, 2, 3}
	s2 := make([]int, 3)
	fmt.Printf("%p, %d, %d\n", s1, cap(s1), len(s1))
	fmt.Printf("%p, %d, %d\n", s2, cap(s2), len(s2))
}

func main() {
	test4()
}
