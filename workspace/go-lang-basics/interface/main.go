package main

import (
	"fmt"
	"os"
)

type Box struct {
	Width int
	Height int
}

type Animal interface {
	Speak() string
}

type Dog struct {

}

func (d Dog)Speak()(string) {
	return "Dog spk"
}

func main() {
	/*
	b := Box{Width:1,Height:4}
	fmt.Println(b)
	fmt.Println(b.Width)

	 */
	var dog Animal = Dog{}

	ans := dog.Speak()
	fmt.Println(ans)


	file, err := os.Open("samle.txt")

	if err != nil {
		fmt.Println("lol")
		fmt.Println(file)

	} else {
		fmt.Println("lol")
		fmt.Println(err)
	}

}
