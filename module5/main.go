package main

import "fmt"

func main() {
	fmt.Println("Pointers Overview")

	// var test *int
	// fmt.Println("Value of pointer", test)

	var t2 = 24

	var ptr = &t2
	*ptr = 26
	fmt.Println(ptr)
	fmt.Println(*ptr)

}
