package main

import "fmt"

func main() {

	switch day := 4; day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thrusday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid")
	}

	var value interface{}
	fmt.Println(value)
	switch value.(type) {
	case bool:
		fmt.Println("value is bool")
	case int:
		fmt.Println("value is int")
	case string:
		fmt.Println("value is string")
	default:
		fmt.Println("okay")
	}
}
