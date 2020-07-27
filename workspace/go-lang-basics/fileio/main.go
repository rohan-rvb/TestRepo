package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("sample.txt")

	if err != nil {
		fmt.Println(file)

	} else {
		fmt.Println("lol")
		fmt.Println(err)
	}
}
