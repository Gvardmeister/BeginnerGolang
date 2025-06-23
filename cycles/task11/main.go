package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	sum := 0
	mul := 1
	result := 0

	for num != 0 {
		result = num % 10
		sum += result
		mul *= result
		num = num / 10
	}
	fmt.Println(sum)
	fmt.Println(mul)
}
