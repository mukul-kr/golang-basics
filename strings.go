package main

import "fmt"

func main() {
	var a string = "Mukul" // variables are immutable so we cannot change part of variable like we were able to do in c string
	b := "hi\n"            // this is what u call a normal string like it does what it is meant for like when you use "\n" it gives new line
	c := `idk\n`           // but here when one use "\n" it doesn't give a new line so it is what it is..
	fmt.Println(a, b, c)
	// let's itterate over a string
	//  as you can see we are using two variables index and chr one takes index of current element and other takes value present at current index

	for index, chr := range a + b + c {
		fmt.Printf("%c  %d\n", chr, index)
	}
	for i := 0; i < len(a); i++ {
		fmt.Printf("%c", a[i])
	}
	xyz_slice := []byte{0x47, 0x6b, 0x65}
	d := string(xyz_slice)
	fmt.Println(d)
	fmt.Println(len(a + b + c))
}
