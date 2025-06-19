package main

import "fmt"

func main() {
	fmt.Println(swap(3, 5))
}

func swap(a int, b int) (int, int) {
	a, b = b, a

	return a, b
}
