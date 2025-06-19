package main

import (
	"fmt"
)

func main() {
	var num1, num2, num3 float64
	fmt.Scan(&num1, &num2, &num3)

	result := (num1 + num2 + num3) / 3
	res := fmt.Sprintf("%.2f", result)

	fmt.Println(res)
}
