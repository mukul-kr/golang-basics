package main

import "fmt"

type interfeace_name interface {
	add() int
	subtraction() int
	multiplication() int
	division() int
}

type two_num struct {
	num1 int
	num2 int
}

func (t two_num) add() int {
	return t.num1 + t.num2
}
func (t two_num) subtraction() int {
	return t.num1 - t.num2
}
func (t two_num) multiplication() int {
	return t.num1 * t.num2
}
func (t two_num) division() int {
	return t.num1 / t.num2
}
func main() {
	var t interfeace_name
	t = two_num{
		num1: 10,
		num2: 5,
	}
	fmt.Println(t.add())
	fmt.Println(t.subtraction())
	fmt.Println(t.multiplication())
	fmt.Println(t.division())

}
