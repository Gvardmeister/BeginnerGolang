package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	result := mapResult(str)

	fmt.Print("map[")
	first := true
	for r, count := range result {
		if !first {
			fmt.Print(" ")
		}
		first = false
		fmt.Printf("%c:%d", r, count)
	}
	fmt.Println("]")
}

func mapResult(str string) map[rune]int {
	newMap := make(map[rune]int)

	for _, value := range str {
		newMap[value]++
	}

	return newMap
}
