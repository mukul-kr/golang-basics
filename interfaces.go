package main

import "fmt"

// creating a type of interface which will calulate all passed function at time of passing it like later we can se that after creating interface we created struct over it so all function were calculated at that time only which is a good thing
type interfeace_name interface {
	add() int
	subtraction() int
	multiplication() int
	division() int
}

// create a basic struct which has two sub types as int
type two_num struct {
	num1 int
	num2 int
}

// 4 function which do addition sub mul div
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
	var t interfeace_name = two_num{
		num1: 10,
		num2: 5,
	}
	fmt.Println(t.add())
	fmt.Println(t.subtraction())
	fmt.Println(t.multiplication())
	fmt.Println(t.division())

	var a1 interface {
	} = 98.09

	myfun(a1)

	var a2 interface {
	} = "GeeksforGeeks"

	myfun(a2)

}

func myfun(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}
