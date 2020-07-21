package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal interface {
	Eat()
	Sleep()
	Breathe()
}

type Address struct {
	house_no int64
	street string
	city string
	pin int64
}

type Employee struct {
	id int64
	name string
	salary float64
	address Address
}

func (emp Employee) Eat() {
	fmt.Println("Eat method")
}

func (emp Employee) Sleep() {
	fmt.Println("Sleep method")
}

func (emp Employee) Breathe() {
	fmt.Println("Breathe method")
}

func main() {
	var emp Employee
	var add Address

	employees := make([]Employee, 0)
	reader := bufio.NewReader(os.Stdin)
	i := 0

	for {
		fmt.Println("Enter employee id")
		id_input, _ := reader.ReadString('\n')
		id, _ := strconv.ParseInt(strings.TrimSpace(id_input),10,32)
		emp.id = id

		fmt.Println("Enter employee name")
		name_input, _ := reader.ReadString('\n')
		name := strings.TrimSpace(name_input)
		emp.name = name

		fmt.Println("Enter employee salary")
		salary_input, _ := reader.ReadString('\n')
		salary, _ := strconv.ParseFloat(strings.TrimSpace(salary_input),64)
		emp.salary = salary

		fmt.Println("Enter city")
		city_input, _ := reader.ReadString('\n')
		city := strings.TrimSpace(city_input)
		add.city = city
		emp.address = add

		employees = append(employees, emp)

		if i < 5 {
			fmt.Println("Do you want to quit (yes/no) ")
			cont, _ := reader.ReadString('\n')
			cont = strings.TrimSpace(cont)

			if cont == "yes" {
				break
			}
		}

		i++
	}

	for _, e := range employees {
		e.Breathe()
		fmt.Println(e)
	}
}