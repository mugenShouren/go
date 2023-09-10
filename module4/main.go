package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time pkg tutorial")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	loc, _ := time.LoadLocation("Asia/Kolkata")
	createdDate := time.Date(1998, time.October, 24, 10, 30, 0, 0, loc)

	fmt.Println("Your Birthdate is: ", createdDate)
}
