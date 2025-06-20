package main

import "fmt"

// Числа Фибоначчи
func main() {
	var num int
	fmt.Scan(&num)
	a, b := 0, 1
	c := 0

	fmt.Print(a, b, " ")

	for i := 3; i <= num; i++ {
		fmt.Print(a+b, " ")

		c = b
		b = a + b
		a = c
	}
}
