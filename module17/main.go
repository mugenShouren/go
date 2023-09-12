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

type circledata struct {
	Name     string   `json:"firstname"`
	Age      int      `json:"age"`
	Company  string   `json:"currentCompany"`
	Password string   `json:"-"`
	Skills   []string `json:"skills,omitempty"`
}

func main() {
	fmt.Println("Playing with JSON")
	// encodeJson()

	jsonData := []byte(`
		{	
			"firstname": "Ayush",   
			"Age": 24,
			"skills": [
					"Python",       
					"go",
					"AWS"
			]
		}	
	`)

	// decodeJson(jsonData)
	decodeJsonAsMap(jsonData)
}

func encodeJson() {
	circle := []circledata{
		{"Ayush", 25, "BR", "1213$%", []string{"Python", "go", "AWS"}},
		{"Shivangi", 24, "BoFA", "1HEH3$(*)", []string{"Python", "SQL", "Quartz"}},
		{"Palash", 24, "CZ", "FJASEF84^)(&())", nil},
	}

	jsonData, err := json.MarshalIndent(circle, "", "\t")
	checkError(err)
	fmt.Printf("%s\n", jsonData)
	fmt.Printf("%T\n", jsonData)

}

func decodeJson(data []byte) {
	jsonValidity := json.Valid(data)

	var returnData circledata

	if jsonValidity {
		fmt.Println("Valid JSON")
		json.Unmarshal(data, &returnData)
		fmt.Printf("%#v\n", returnData)
	} else {
		fmt.Println("Invalid JSON")
	}
}

func decodeJsonAsMap(data []byte) {

	jsonValidity := json.Valid(data)
	var circleMap map[string]interface{}

	if jsonValidity {
		fmt.Println("Valid JSON")
		json.Unmarshal(data, &circleMap)
		fmt.Printf("%#v\n", circleMap)
	} else {
		fmt.Println("Invalid JSON")
	}

	for k, v := range circleMap {
		fmt.Printf("The key is: %v and value is: %v and its type is: %T\n", k, v, v)
	}

}
