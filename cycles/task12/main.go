package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Scan(&num)

	max := 0
	result := 0

	for num != 0 {
		result = num % 10

		if result > max {
			max = result
		}

		num = num / 10
	}
	fmt.Println(max)
}
