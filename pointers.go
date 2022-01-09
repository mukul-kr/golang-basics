package main

import "fmt"

func main() {
	// Creating a pointer
	var a = 5
	var ptra *int = &a // creating a pointer
	fmt.Println(*ptra)

	// short hand declaration
	b := 10
	ptrb := &b
	fmt.Println(*ptrb)

	var ptrn *int
	fmt.Println(ptrn) // while dereferencing this nill pointer will generate error

}
