package main

import (
	"fmt"
	"sort"
)

func main() {
	//array
	var num  = [3]int{1, 2, 3}
	arr := []int{10, 40}
	fmt.Println(num)
	fmt.Println(arr)


	//slice
	arr = append(arr, 5)
	fmt.Println(arr)

	sort.Ints(arr)
	//sort.Reverse(arr)
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
