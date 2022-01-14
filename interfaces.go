package main

import "fmt"

// creating a type of interface which will run in following given funcitons i guess...
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

type st1 struct {
	ta1 int
	na1 string
}
type st2 struct {
	ta2 int
	na2 string
}

var inter2 interface {
	display()
}

func (s st2) display() {
	fmt.Println(s.ta2)
}

var inter1 interface {
	display()
}

func (s st1) display() {
	fmt.Println(s.na1)
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
	} = 21.01
	myfun(a1) // a1 interface holds value 21.01 and my fun was called through a1 so will execute accordingly

	var a2 interface {
	} = "hi bitxh"
	myfun(a2) // a2 interface holds value "hi bitxh" and my fun was called through a2 so will execute accordingly

	// polymorphism through interfaces
	s := st1{
		ta1: 1,
		na1: "hi",
	}
	s.display()
	s2 := st1{
		ta1: 2,
		na1: "hello",
	}
	s2.display()

}

func myfun(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}
