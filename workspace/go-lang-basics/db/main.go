package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "hello file test"

	file, err := os.Create("./sample.txt")
	checkError(err)
	n, err := io.WriteString(file, content)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
