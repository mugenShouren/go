package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Hey yo Ayush"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating:")

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of the rating %T", input)
}
