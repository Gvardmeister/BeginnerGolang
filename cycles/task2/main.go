package main

import (
	"fmt"
)

func main() {
	var number1, number2 int
	fmt.Scan(&number1, &number2)

	for i := number1; i <= number2; i++ {
		fmt.Print(i*i*i, " ")
	}
}
