package main

import "fmt"

func main() {
	// If statement is same as all other language
	if 5 < 6 {
		fmt.Println("👀 👀")
	}
	if 6 < 5 {
		fmt.Println("👀")
	}
	// but I have discovered one thing that If statement only takes boolean values as argument and not integer values...
	// so unlike pyhton or c++ in which we can give both true or any other non-zero number for if statement to work and false or zero to if block to not execute .
	// But this isn't possible here 😢

	//
	// Nested if and if else statements
	if 4 < 7 {
		fmt.Println(1)
		if 5 < 9 {
			fmt.Println(2)
		}
	} else if 6 < 9 { // else if statement should start just after code block of if statement
		fmt.Println(3)
	} else { // else statement should start just after code block of else if or if statement
		fmt.Println(4)
	}

}
