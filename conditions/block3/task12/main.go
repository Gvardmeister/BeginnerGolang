package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num > 100 {
		fmt.Println("Много")
	} else {
		fmt.Println("Норм")
	}
}
