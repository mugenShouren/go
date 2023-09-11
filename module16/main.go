package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("Making CRUD requests")
	const baseUrl string = "http://localhost:8000"
	getRequest(baseUrl, "/get")

	payload := strings.NewReader(`
	{
		"name" : "ayush",
		"age" : 24
	}
	`)
	postRequest(baseUrl, "/post", payload)

	payload2 := url.Values{}
	payload2.Add("FirstName", "Ayush")
	payload2.Add("LastName", "Khare")
	payload2.Add("Email", "khaayush@google.com")
	postFormRequest(baseUrl, "/postform", payload2)
}

func postFormRequest(baseUrl string, path string, payload url.Values) {

	resp, err := http.PostForm(baseUrl+path, payload)
	checkError(err)
	defer resp.Body.Close()
	var responseString strings.Builder
	body, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(body)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Printf("The reponse returned is: %v\n", string(body))

}

func postRequest(baseUrl string, path string, payload *strings.Reader) {

	resp, err := http.Post(baseUrl+path, "application/json", payload)
	checkError(err)
	defer resp.Body.Close()
	var responseString strings.Builder
	body, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(body)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Printf("The reponse returned is: %v\n", string(body))

}

func getRequest(baseUrl string, path string) {
	resp, err := http.Get(baseUrl + path)
	checkError(err)
	defer resp.Body.Close()
	var responseString strings.Builder
	body, _ := ioutil.ReadAll(resp.Body)
	byteCount, _ := responseString.Write(body)
	fmt.Println("Byte count is: ", byteCount)
	fmt.Printf("The reponse returned is: %v\n", string(body))

	fmt.Println("String from second method: ", responseString.String())
}
