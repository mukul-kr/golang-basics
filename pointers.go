package main

import "fmt"

func goodfunction(ptr *int) {
	*ptr = *ptr + 143
}

func main() {
	// Creating a pointer
	var a = 5
	var ptra *int = &a // creating a pointer
	fmt.Println(*ptra) // dereferencing a pointer to get value stored in that memory location

	// short hand declaration
	b := 10
	ptrb := &b
	fmt.Println(*ptrb)

	var ptrn *int
	fmt.Println(ptrn) // while dereferencing this nill pointer will generate error

	goodfunction(ptra)
	fmt.Println(a)

	goodfunction(&b)
	fmt.Println(b)
}
