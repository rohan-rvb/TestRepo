package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://google.com"

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	//fmt.Println(resp.Body)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))

}
