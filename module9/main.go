package main

import "fmt"

func main() {
	fmt.Println("Conditional statements")

	val := 51

	var result string

	if val < 20 {
		result = "Ocassional User"
	} else if val > 50 {
		result = "Addict User"
	} else {
		result = "Regular User"
	}

	fmt.Println(result)

	if num1, num2 := 5, 4; num1+num2 > 10 {
		fmt.Println("Numbers are greater than 10")
	} else {
		fmt.Println("Numbers are Not greater than 10")
	}
}
