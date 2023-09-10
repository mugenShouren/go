package main

import "fmt"

func main() {
	fmt.Println("first print")
	defer fmt.Println("second print ")
	fmt.Println("third print")
	defer fmt.Println("fourth print ")
	testDefer()

	defer fmt.Println("\nfifth print")

}

func testDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
