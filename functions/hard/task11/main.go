package main

import "fmt"

func main() {
	var x, y int
	x = 3
	y = 4

	fmt.Println(square(x) + square(y))
}

func square(n int) int {
	return n * n
}
