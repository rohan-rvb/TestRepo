package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("in main")
	go fn()
	time.Sleep(500)
}

func fn() {
	fmt.Println("in routine")
}
