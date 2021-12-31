package main

import "fmt"

type car struct { // a normal struct
	model string
	color string
	seats int
}

type data int // only 1 variable that is int so basically we are storing int here

func (item data) mul2() data {
	return item * 2
}

func (item car) show() {
	fmt.Println(item.model)
	fmt.Println(item.color)
	fmt.Println(item.seats)
}

func main() {
	c1 := car{
		model: "m1",
		color: "blue",
		seats: 7,
	}
	c1.show()
	d1 := data(5)
	fmt.Println(d1.mul2())
}
