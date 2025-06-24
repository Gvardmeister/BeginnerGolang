package main

import "fmt"

func main() {
	arr := []int{5, 3, 8, 6}

	for key, value := range arr {
		if value == 8 {
			fmt.Println(key)
		}
	}
}
