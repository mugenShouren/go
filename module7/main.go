package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	langs := make(map[string]string)

	langs["JS"] = "2000"
	langs["C++"] = "1970"
	langs["PY"] = "1990"
	langs["JAVA"] = "1980"

	fmt.Println("List of keys", langs)
	fmt.Println("JS value:", langs["JS"])

	delete(langs, "JAVA")
	fmt.Println("List of keys", langs)

	for key, value := range langs {
		fmt.Printf("Key: %v Value: %v\n", key, value)
	}

}
