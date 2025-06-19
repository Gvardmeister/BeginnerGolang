package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	if num%3 == 0 {
		fmt.Println("Делится")
	}
}
