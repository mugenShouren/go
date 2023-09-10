package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("Working with functions")

	func1()

	result := addition(3, 5)

	fmt.Println("Result is: ", result)

	proResult, _ := proAddition(1, 2, 3, 4, 5, 6)

	fmt.Println("PRO Result is: ", proResult)

}

func addition(v1 int, v2 int) int {
	return v1 + v2
}

func proAddition(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, strconv.Itoa(total)
}

func func1() {
	fmt.Println("Basic call")
}
