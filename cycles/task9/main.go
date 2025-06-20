package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	sum := 0
	result := 0
	for num != 0 {
		result = num % 10

		if result%2 == 0 {
			sum += result
		}

		num = num / 10
	}
	fmt.Println(sum)
}
