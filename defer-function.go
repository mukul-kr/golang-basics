package main

import "fmt"

func f1(n int) {
	fmt.Println(n)
}
func main() {
	f1(1)
	f1(2)
	f1(3)
	f1(4)

	defer f1(10)
	defer f1(20)
	defer f1(30)
	defer f1(40)
}
