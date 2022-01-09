package main

import "fmt"

type Employee struct {
	name  string
	empid int
}

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

	// passing address to function through pointers
	goodfunction(ptra)
	fmt.Println(a)

	goodfunction(&b)
	fmt.Println(b)

	// pointer to a struct
	emp := Employee{"Mukul", 143}
	ptrs := &emp
	fmt.Println((*ptrs).name)

	d := "doyouloveme-givemeahint-pls"
	ptrd := &d
	p_ptrd := &ptrd
	fmt.Println(**p_ptrd)
}
