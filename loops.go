package main

import "fmt"

func main() {
	// simple for loop
	for i := 0; i < 3; i++ {
		fmt.Print("a")
	}
	j := 0
	// for loop as infinite loop
	for {
		fmt.Print("b")
		j++
		if j == 2 {
			break
		}
	}
	// for loop as while loop
	for j < 5 {
		fmt.Print("c")
		j++
	}
	// for loop to itterate over an item as we do in python

	str_array := []string{"a", "b", "c", "d"}
	for i, j := range str_array {
		fmt.Println(i, j)
	}

	// for loop to itterate over a string
	str := "Mukul"
	for index, chr := range str {
		fmt.Println(index, chr) // Here ascii code is printed instead of actual value
	}

	// For loop to itterate over map value
	map_example := map[int]string{
		1: "a",
		2: "b",
		3: "c", // One thing to note is that even when there is no value after 3:"c" we are still required to put a comma after that or compiler will generate error
	}
	for key, value := range map_example {
		fmt.Println(key, value)
	}

	// For channel
	chnl := make(chan int)
	go func() {
		chnl <- 100
		chnl <- 1000
		chnl <- 10000
		chnl <- 100000
		close(chnl)
	}()
	for i := range chnl {
		fmt.Println(i)
	}
}
