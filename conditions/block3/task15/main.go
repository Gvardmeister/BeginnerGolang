package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)

	if age < 7 {
		fmt.Println("Ты малыш")
	}
}
