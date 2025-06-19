package main

import "fmt"

func main() {
	double := makeMultiplier(2)

	fmt.Println(double(5))
}

func makeMultiplier(factor int) func(int) int {
	return func(i int) int { return factor * i }
}
