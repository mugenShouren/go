package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")

	ayush := User{"Ayush", "khaayush@google.com", true, 24}

	fmt.Println(ayush)

	fmt.Printf("Ayush's details are %+v\n", ayush)

	ayush.getStatus()

	ayush.setEmail("khareayush@google.com")

	ayush.getEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) getStatus() {
	fmt.Println("The status of user is ", u.Status)
}

func (u *User) setEmail(newEmail string) {
	u.Email = newEmail
}

func (u User) getEmail() {
	fmt.Println("The Email of user is ", u.Email)
}
