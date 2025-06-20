package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	result := 0
	for num != 0 {
		result = num % 10
		num = num / 10

		fmt.Println(result)
	}
}
