package main

import "fmt"

func main() {
	fmt.Println(double(10))
}

func double(n int) int {
	return n * 2
}
