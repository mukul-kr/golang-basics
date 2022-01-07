package main

import "fmt"

func main() {
	// array in go
	// arrays in go behave in same way as c like it has fixed size and it can only hold values of same datatype.
	var a [5]int
	fmt.Println(a) //default initialises all value in array by zero
	a[0] = 1
	a[1] = 3
	a[4] = 10
	fmt.Println(a)
	var b [3]string
	fmt.Println(b)
	b[0] = "I"
	b[1] = "Love"
	b[2] = "Sonali"
	fmt.Println(b)

}
