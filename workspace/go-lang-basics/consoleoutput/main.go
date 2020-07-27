package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("ft")

	str1 := "This is string 1"

	out, _ := fmt.Println(str1)
	fmt.Println(out)

	str2 := "This is second string"
	str3 := "hello there"
	aNumber := 33
	isTrue := true
	// if you dont want the return parameter the put _
	strLength, _ := fmt.Println(str1, str2, str3);
	//if err == nil {
	// fmt.Println("The length of the string: ", strLength);
	//}
	fmt.Println("The length of the string: ", strLength)
	fmt.Printf("Value of aNumber %v\n", aNumber)
	fmt.Printf("Value of Boolean %v \n", isTrue)
	fmt.Printf("Value of aNumber as float: %.2f \n", float64(aNumber))

	fmt.Printf("data types %T %T %T \n", str1, aNumber, isTrue)

	fmt.Sprint("hello")
	//stt := fmt.Sprint("hello")
	//fmt.Println(stt)

	reader := bufio.NewReader(os.Stdin)
	//r, _ := reader.ReadString('\n')

	sal, _ := reader.ReadString('\n')
	sal = strings.TrimSpace(sal)
	fmt.Println("sa",sal)
	fsal, _ := strconv.ParseFloat(sal,32)

	month, _ := reader.ReadString('\n')
	month = strings.TrimSpace(month)
	a, _ := strconv.ParseInt(month,10,32)

	final := float64(a)*fsal
	fmt.Println("salary is ", fsal, final)


}
