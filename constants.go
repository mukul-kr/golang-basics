package main

import (
	"fmt"
	"reflect"
)

func main() {
	const PI = 3.14
	const hi = "🙂"
	const correct = true
	fmt.Println(PI, hi, correct)
	var x1 = 85
	fmt.Println(x1, reflect.TypeOf(x1))
	var x2 = 0213
	fmt.Println(x2, reflect.TypeOf(x2))
	var x3 = 0x4b
	fmt.Println(x3, reflect.TypeOf(x3))
	x4 := 30 // 30 u giving error but it shouldn't have given error will check later why this happened.
	fmt.Println(x4, reflect.TypeOf(x4))
	var x5 = 30 // 30 l is also giving error but it shouldnt have thrown that error.
	fmt.Println(x5, reflect.TypeOf(x5))
}
