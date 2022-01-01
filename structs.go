package main

import "fmt"

// just a normal struct
type Address struct {
	name    string
	city    string
	pincode int
}

// nested struct
type student struct {
	name string
	add  Address
}

func main() {

	var obj Address = Address{
		name:    "Mukul",
		city:    "Guna",
		pincode: 521463,
	}
	fmt.Println(obj)
	fmt.Println(obj.name) // struct calling through objects
	ptr := &obj
	fmt.Println((*ptr).city) //struct calling through pointer
	fmt.Println(ptr.city)
	new_obj := student{
		name: "mukul",
		add: Address{
			name:    "hi",
			city:    "fhd",
			pincode: 153454,
		},
	}
	fmt.Println(new_obj)          // nested struct
	fmt.Println(new_obj.add.city) //nested struct member calling

	// anonymous struct
	st := struct {
		name string
		val  int
	}{
		name: "Mukul",
		val:  5,
	}
	fmt.Println(st)
}
