package main

import (
	"fmt"
	"math/big"
	"time"
)

func main() {
	//i1, i2, i3 := 1, 2, 3
	//sum := i1 + i2 + i3

	var b1, b2, b3, bg big.Float
	b1.SetFloat64(1.1)
	b2.SetFloat64(1.2)
	b3.SetFloat64(1.3)

	bg.Add(&b1, &b2).Add(&bg, &b3)
	fmt.Printf("bg val %f \n",&bg)

	t := time.Now()
	fmt.Println(t)

	t1 := t.AddDate(0,0,1)
	fmt.Println(t1)

	fmt.Println(t.Format("Monday"))
}
