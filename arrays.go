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

	// Multi-dimensional array
	arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(arr)

	// "if ellipsis" "..." here length is determined by initialized elements
	arr1 := [...]int{1, 2, 3, 4}
	fmt.Println(len(arr1))
	// above we have used len function which calculates the lenght of passed array

	// traversing over an array of elements
	for index, val := range arr1 {
		fmt.Println(index, val)
	}

	// In go array is of value type and not reference type so changes made to a copy is not refelected back in original array
	arr2 := arr1
	arr2[0] = 10
	fmt.Println("arr2", arr2)
	fmt.Println("arr1", arr1)
	arr3 := &arr1
	arr3[0] = 100 //copying via referencing or better say it as that we have stored address of arr1 in arr3 and we will be making changes in arr1 though its address
	fmt.Println("arr2", *arr3)
	fmt.Println("arr1", arr1)

	// If two array are of same type then we can eaislt compare it....
	fmt.Println(arr1 == arr2)
	fmt.Println(arr1 == *arr3)

	// Passing array into function
	fmt.Println(sum(arr1))
}
func sum(arr [4]int) (result int) {
	for _, item := range arr {
		result += item
	}
	return
}
