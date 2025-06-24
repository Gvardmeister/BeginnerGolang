package main

import "fmt"

func main() {
	var slice [5]int = [5]int{3, 5, 2, 7, 1}

	sum := 0
	for _, value := range slice {
		sum += value
	}
	fmt.Println(sum)
}
