package main

import "fmt"

func main() {
	m := make(map[int]int)
	m[1] = 2
	m[2] = 3

	key := 5
	m[key] = 100
	a, b := m[key]
	fmt.Println(a,b)
	fmt.Println(m)
	var st, ca []int

	for p, q := range m {
		fmt.Println(p, q)
		st = append(st, p)
		ca = append(ca, q)
	}
	fmt.Println(st)
	fmt.Println(ca)
}