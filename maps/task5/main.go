package main

import "fmt"

func main() {
	mapFirst := map[string]int{
		"a": 1,
		"b": 2,
	}
	mapSecond := map[string]int{
		"b": 2,
		"a": 1,
	}

	fmt.Println(twoMaps(mapFirst, mapSecond))
}

func twoMaps(mapFirst, mapSecond map[string]int) bool {
	if len(mapFirst) != len(mapSecond) {
		return false
	}

	for k, valueFirst := range mapFirst {
		valueSecond, ok := mapSecond[k]

		if !ok || valueFirst != valueSecond {
			return false
		}
	}

	return true
}
