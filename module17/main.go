package main

import (
	"encoding/json"
	"fmt"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type somedata struct {
	Name    string   `json:"firstname"`
	Age     int      `json:"-"`
	company string   `json:"currentCompany"`
	Skills  []string `json:"skills,omitempty"`
}

func main() {
	fmt.Println("Playing with JSON")

	encodeJson()

}

func encodeJson() {
	circle := []somedata{
		{"Ayush", 24, "BR", []string{"Python", "go", "AWS"}},
		{"Shivangi", 24, "BR", []string{"Python", "SQL", "Quartz"}},
		{"Palash", 24, "BR", nil},
	}

	jsonData, err := json.MarshalIndent(circle, "", "\t")
	checkError(err)
	fmt.Printf("%s\n", jsonData)

}
