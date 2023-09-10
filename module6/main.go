package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println("Playing with Slices")

	var fL = []string{"Mango", "banana"}
	var arr = [2]string{}
	fmt.Printf("Type of lists: %T %T\n", fL, arr)

	fL = append(fL, "apple", "pineapple")
	fmt.Println(fL)

	fL = append(fL[1:3])
	fmt.Println(fL)

	highScores := make([]int, 4)

	highScores[0] = 349
	highScores[1] = 783
	highScores[2] = 567
	highScores[3] = 155

	highScores = append(highScores, 222, 444, 555)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	var strings = []string{"1", "2", "3"}
	index := 1
	strings = append(strings[:index], strings[index+1:]...)

	fmt.Println(strings)

}
