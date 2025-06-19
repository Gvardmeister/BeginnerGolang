package main

import "fmt"

func main() {
	var cost int
	fmt.Scan(&cost)

	fmt.Println(cost - (cost * 13 / 100))
}
