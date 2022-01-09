package main

import (
	"fmt"
	"sort"
)

func main() {
	// creating a slice
	var slc0 = []string{"I", "don't", "want", "to", "live"}
	fmt.Println(slc0)

	// short hand notation slice
	slc := []int{1, 2, 3}
	fmt.Println(slc)

	// creating slice from array
	// It basically works like we used to slice an list in python just that we are storing it in a variable
	arr := [5]int{1, 2, 3, 4, 5}
	slc1 := arr[1:4]
	fmt.Println(slc1)
	fmt.Println(len(slc1))
	fmt.Println(cap(slc1))

	// creating slice through make
	var slc2 = make([]int, 3, 7)
	var slc3 = make([]int, 7)
	fmt.Println(slc2)
	fmt.Println(slc3)

	// itterating over a slice
	for _, item := range slc1 {
		fmt.Println(item)
	}

	// modifying slice
	slc1[1] = 100
	fmt.Println(arr)
	fmt.Println(slc1)

	// zero value slcie
	var slc4 []int
	fmt.Println(slc4, len(slc4), cap(slc4))

	// 2dimensional slice
	slc5 := [][]int{{1, 2, 3}, {1, 2}}
	fmt.Println(slc5)

	// sorting of slice
	slc6 := []int{9, 8, 7, 4, 6, 8, 5}
	fmt.Println(slc6)
	sort.Ints(slc6)
	fmt.Println(slc6)

}
