package main

import "fmt"

func main() {
	var array [8]int = [8]int{2, 7, 8, 3, 4, 9, 6, 1}

	sumChet := 0
	for _, value := range array {
		if value%2 == 0 {
			sumChet++
		}
	}
	fmt.Println(sumChet)
}
