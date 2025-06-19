package main

import "fmt"

func main() {
	fmt.Println(sum(5, 10))
	fmt.Println(sum(7, 3))
	fmt.Println(sum(9, 100))
}

func sum(a, b int) int {
	return a + b
}
