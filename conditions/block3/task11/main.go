package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num == 0 {
		fmt.Println("Число равно 0")
	} else {
		fmt.Println("Число НЕ равно 0")
	}
}
