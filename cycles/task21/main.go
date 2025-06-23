package main

import "fmt"

func main() {
	const start, end = 32, 126
	count := 0

	for i := start; i <= end; i++ {
		fmt.Printf("%3d:%c  ", i, i)
		count++
		if count%5 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
