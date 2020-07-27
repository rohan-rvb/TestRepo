package main

import "fmt"

func main() {
	res := Add(1, 2)
	fmt.Println(res)
}

func Add(a int, b int) int {
	return a+b
}