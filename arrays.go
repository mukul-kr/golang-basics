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
	// I was happy after meeting you today so wanted to tell someone but didn't find anyone whom I can tell what I felt so writing in this code
	// So if you are reading this ignore it
	// Just expressing my inner feelings to my code
	fmt.Println(b)
	// we can also assign values at the time of creating vaariables
	var c [3]string = [3]string{"a", "b", "c"}
	fmt.Println(c)
	// short hand declaration for array
	d := [5]int{1, 2, 3, 4, 5}
	e := [5]int{}
	fmt.Println(d)
	fmt.Println(e)

}
