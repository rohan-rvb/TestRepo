package main

import (
	"fmt"
	"net/http"
)

func main() {
	urls := []string {
		"https://golang.org/",
		"https://jsonplaceholder.typicode.com/posts",
	}

	for _, url := range urls {
		//returnType(url)
		go returnType(url)
	}
}

func returnType (url string) {
	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()

	ctype := resp.Header.Get("content-type")
	fmt.Printf("%s -> %s \n", url, ctype)
}

func checkError(err error) {
	if err !=nil {
		panic(err)
	}
}

