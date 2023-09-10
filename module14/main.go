package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Playig with web")

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response %T\n", resp)

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	content := string(data)

	fmt.Println(content)
}
