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

func fun1(channel1 chan string) {
	time.Sleep(1 * time.Second)
	channel1 <- "Hi 🙂"
}
func fun2(channel2 chan string) {
	time.Sleep(2 * time.Second)
	channel2 <- "hello 🤣"
}
func main() {
	go display("hi")
	display("bye")

	// go display("hi")
	time.Sleep(1 * time.Second)
	display("bye")

	ch1 := make(chan string)
	ch2 := make(chan string)

	go fun1(ch1)
	go fun2(ch2)

	select {
	case option1 := <-ch1:
		fmt.Println(option1)
	case option2 := <-ch2:
		fmt.Println(option2)
	}

}
