package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	new := 0
	result := 0
	for num != 0 {
		result = num % 10
		new = new*10 + result
		num = num / 10
	}
	fmt.Println(new)
}
