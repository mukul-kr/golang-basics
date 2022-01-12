package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
}

func main() {
	go display("hi")
	display("bye")

	// go display("hi")
	time.Sleep(1 * time.Second)
	display("bye")
}
