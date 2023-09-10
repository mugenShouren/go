package main

import "fmt"

const url string = "https://lco.dev:3000/learn?course=golang&user=ayush"

func main() {
	fmt.Println("Playing with URLs")

	fmt.Println(url)
}
