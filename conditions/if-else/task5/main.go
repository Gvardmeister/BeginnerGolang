package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num > 0 {
		fmt.Println("Положительное число")
	} else {
		fmt.Println("Отрицательное число")
	}
}
