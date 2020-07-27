package main

import "fmt"

func main() {
	/*
	go DoMyWork("roha")
	go DoMyWork("pro")

	time.Sleep(time.Second*3)

	 */

	ch := make(chan string)
	//ch <-1

	go greet(ch)
	//go rty(ch)
	ch <- "Rohan"
	//ch <- "bha"
	close(ch)
	fmt.Println("end")
}

func rty(c chan string) {
	fmt.Println("hyhy"+ <-c)
}

func greet(c chan string) {
	fmt.Println("Hello "+<-c)
}

/*
func DoMyWork(name string) {
	for i :=1 ; true ; i++ {
		fmt.Println("work of ", name, "exec ", i)
	}
}

 */
