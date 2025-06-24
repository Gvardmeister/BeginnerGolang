package main

import "fmt"

func main() {
	var array [10]int = [10]int{1, 4, 6, 8, 3, 9, 2, 0, 5, 7}

	maxValue := array[0]
	for _, value := range array {
		if value > maxValue {
			maxValue = value
		}
	}
	fmt.Println(maxValue)
}
