package main

import "fmt"

func main() {
	mapFirst := map[string]string{
		"cat": "кот",
		"dog": "собака",
	}

	fmt.Println(reversed(mapFirst))
}

func reversed(mapFirst map[string]string) map[string]string {
	newMap := make(map[string]string)

	for key, value := range mapFirst {
		newMap[value] = key
	}

	return newMap
}
