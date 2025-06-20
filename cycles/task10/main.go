package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	even := 0
	uneven := 0
	dev := 0

	for num != 0 {
		dev = num % 10

		if dev%2 == 0 {
			even++
		} else {
			uneven++
		}

		num = num / 10
	}
	fmt.Println(even)
	fmt.Println(uneven)
}
