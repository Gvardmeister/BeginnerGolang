package main

import "fmt"

func main() {
	arr := []int{2, 4, 6, 8}

	sum := 0
	for _, value := range arr {
		sum += value
	}
	fmt.Println(sum / 4)
}
