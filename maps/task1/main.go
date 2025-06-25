package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	fmt.Println(mapResult(num))
}

func mapResult(num int) map[int]int {
	numMaps := make(map[int]int)

	if num == 0 {
		numMaps[0] = 1

		return numMaps
	}

	for num > 0 {
		digit := num % 10
		numMaps[digit]++
		num = num / 10
	}

	return numMaps
}
