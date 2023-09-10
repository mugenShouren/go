package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	ayush := User{"Ayush", "khaayush@google.com", true, 24}

	fmt.Println(ayush)

	fmt.Printf("Ayush's details are %+v", ayush)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
