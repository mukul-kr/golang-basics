package main

import "fmt"

//lets create a function that adds two number pretty simple

func add_two(a, b int) int {
	y := a + b
	return y
}

//like c or c++ value passed in a function is passed by value so to create a function which swaps value we have to pass it by reference
func my_swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

// there is somethin in go known as varidic function
// it is known so because it allows unknown ammount of variables to be passed in function
func add_n(n ...int) int {
	ans := 0
	for _, j := range n {
		ans += j
	}
	return ans
}

func func_passed(passed_func func(a ...int) int, a ...int) {
	fmt.Println(passed_func(a...))
}

func func_return() func(a, b int) int {
	fn := func(a, b int) int {
		return (a * 2) / b
	}
	return fn
}

func main() {
	a := 1
	b := 2
	x := add_two(a, b)
	fmt.Println(a, b, x)
	my_swap(&a, &b)
	fmt.Println(a, b)
	fmt.Println(add_n(1, 2, 3, 4, 5, 6))

	my_list := []int{2, 3, 4, 5, 6}
	fmt.Println(add_n(my_list...)) // I really like this you ask why ?
	// because u write function that works for solo variable and also for arrays
	// this discovery made me more attracted toward go🥰😍😘

	func() {
		fmt.Println("hello")
	}()

	f1 := func() {
		fmt.Println("arigatou")
	}
	f1()
	func_passed(add_n, 1, 2, 3, 4)

	fr := func_return()
	fmt.Println(fr(5, 10))

}
