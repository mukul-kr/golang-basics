package main

import "fmt"

func main() {
	x := 52
	y := 10
	// basic operator present in all language
	result := x + y
	fmt.Println(result)
	result = x - y
	fmt.Println(result)
	result = x * y
	fmt.Println(result)

	result3 := x / y // Note that even on declaration operator division operator only gives int value and not float value
	// there is one more thing to note that like python in which if use '//' for division int value is returned same is done through '/' operator but there is no operator that I know which can calculate division in float value💁‍♂️😵.
	fmt.Println(result3)

	result = x % y
	fmt.Println(result)
	//
	//realtional operators is same as python 😁.
	result1 := x < y
	fmt.Println(result1)
	result1 = x <= y
	fmt.Println(result1)
	result1 = x > y
	fmt.Println(result1)
	result1 = x >= y
	fmt.Println(result1)
	result1 = x == y
	fmt.Println(result1)
	result1 = x != y
	fmt.Println(result1)
	//
	// Logical Opeartors 😶
	// Not same as python but same as c/c++/java so doen't matter
	if x != y && x <= y {
		fmt.Println("hi")
	} //since second condition i.e. x <= y is not true so if block was not executed..
	if x != y || x <= y {
		fmt.Println("hemlo")
	}
	if !(x == y) {
		fmt.Println("I am looking for change")
	}
	//
	// Bitwise operator 😵😵😵
	// tbh i know how this stuff works but can't get my brain to apply it in real problems😢
	result = x & y
	fmt.Println(result)
	result = x | y
	fmt.Println(result)
	result = x ^ y
	fmt.Println(result)
	result = x << 1
	fmt.Println(result)
	result = x >> 1
	fmt.Println(result)
	result = x &^ 1
	fmt.Println(result)

	//
	// Assignment operator
	// Nothing new here
	x = y
	fmt.Println(x)
	x += y
	fmt.Println(x)
	x -= y
	fmt.Println(x)
	x *= y
	fmt.Println(x)
	x /= y
	fmt.Println(x)
	x %= y
	fmt.Println(x)

	//
	// few other operator
	ptr := &x
	fmt.Println(*ptr)
	fmt.Println(ptr) // clearly visible that it returns address in which variable in
	// lets try changing value of x through ptr
	*ptr = 232
	fmt.Println(x)
	// the end
}
