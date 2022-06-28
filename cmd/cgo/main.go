package main

//#include <stdio.h>
//void cfunc(){
//	printf("Hello, world!\n");
//}
import "C"

func main() {
	C.cfunc()
}
