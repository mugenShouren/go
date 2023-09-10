package main

import (
	"fmt"
	"net/url"
)

const theUrl string = "https://lco.dev:3000/learn?course=golang&user=ayush"

func main() {
	fmt.Println("Playing with URLs")

	fmt.Println(theUrl)

	res, _ := url.Parse(theUrl)

	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.RawQuery)
	fmt.Println(res.Port())

	qparams := res.Query()

	fmt.Printf("Type of params %T\n", qparams)

	fmt.Printf("%T", qparams["course"])

	urlConstructor := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "course=go",
	}
	fmt.Println("Url Constructed", urlConstructor.String())
}
