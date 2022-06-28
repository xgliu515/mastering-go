package main

import "fmt"

func test1() {
	s := "hello"
	fmt.Println(s[1:2])

	a := make([]int, 5)
	fmt.Printf("%T, %d, %p\n", a, len(a), a)
	fmt.Printf("%T, %d, %p\n", a, len(a), &a)
	for _, s := range a {
		fmt.Println(s)
	}
	a = append(a, 3)
	fmt.Printf("%T, %d, %p\n", a, len(a), a)
}

func test2() {
	// Slice will double the cap when append ???
	const N = 10
	for i := 0; i < N; i++ {
		a := make([]int, i)
		fmt.Printf("Cap %d, Len %d\n", cap(a), len(a))
		a = append(a, i)
		fmt.Printf("Cap %d, Len %d\n", cap(a), len(a))
		fmt.Println()
	}
}

func test3() {
	// Slice of byte will print all the visible char ???
	s := make([]byte, 5)
	fmt.Printf("Cap %d, Len %d\n", cap(s), len(s))
	s[0] = 'c'
	s[1] = 0
	s[2] = 'd'

	fmt.Printf("%s\n", s)
	fmt.Printf("%d\n", s[3])
	fmt.Printf("Cap %d, Len %d\n", cap(s), len(s))
}

func test4() {
	a6 := []int{1, 2, 3, 4, 5, 6}
	a4 := []int{7, 8, 9, 0}

	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)

	copy(a6, a4)

	fmt.Println("a6:", a6)
	fmt.Println("a4:", a4)
}

func main() {
	test4()
}
