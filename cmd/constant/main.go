package main

const gc = 1

func test1() {
	const c = 1

	// Const no address
	//fmt.Printf("%p - %p\n", &gc, &c)
}

func main() {
	test1()
}
