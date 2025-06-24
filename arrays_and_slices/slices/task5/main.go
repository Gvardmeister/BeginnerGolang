package main

import "fmt"

func main() {
	sli := []int{1, 2, 3, 4, 5, 6}

	fmt.Println(check(sli))
}

func check(sl []int) []int {
	newSlice := []int{}

	for _, value := range sl {
		if value%2 == 0 {
			newSlice = append(newSlice, value)
		}
	}

	return newSlice
}
