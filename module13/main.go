package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Playing with files")
	content := "Yeaaaah I can write to files FFF"

	file, _ := os.Create("./test.txt")
	defer file.Close()

	io.WriteString(file, content)

	fileReader("./test.txt")

}

func fileReader(filepath string) {
	databyte, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text in file is", string(databyte))
}
