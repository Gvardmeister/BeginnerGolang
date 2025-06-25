package main

import "fmt"

func main() {
	mapCheck := map[string]int{
		"x": 10,
		"y": 20,
	}

	fmt.Println(check(mapCheck, 20))
}

func check(mC map[string]int, target int) string {
	for key, value := range mC {
		if value == target {
			return key
		}
	}

	return "Not found"
}
