package main

import "fmt"

func main() {
	//
	a := 5        // declaration + inititalisation with automatic type assignment
	var b int = 5 // declaration but with manual type assignment
	fmt.Println(a, b)

	var variable int = 10
	var variable2 float32 = 10.01
	fmt.Println(float32(variable) + variable2)
	//

	// above i have discovered a way to add mismatched type variables
	// lets try it on few more types and see what will be result

	//
	x := 10.05
	y := 20
	z := int(x) + y
	fmt.Println(z)
	//

}
