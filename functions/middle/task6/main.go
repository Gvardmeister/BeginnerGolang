package main

import "fmt"

func main() {
	fmt.Println(isEven(10))
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}

	return false
}
