package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	factor := 1
	for i := 1; i <= num; i++ {
		factor *= i
	}
	fmt.Println(factor)
}
