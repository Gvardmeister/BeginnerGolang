package main

import "fmt"

func main() {
	const size = 10

	fmt.Printf("     ")
	for i := 1; i <= size; i++ {
		fmt.Printf("%4d", i)
	}
	fmt.Println()

	fmt.Println("    ", "-------------------------------")

	for i := 1; i <= size; i++ {
		fmt.Printf("%4d|", i)
		for j := 1; j <= size; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}
