package main

import "fmt"

func main() {
	fmt.Println(max(10, 5))
	fmt.Println(max(1, 2))
	fmt.Println(max(3, 3))
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
