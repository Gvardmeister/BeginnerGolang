package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Scan(&number)

	for i := 1; i*i <= number; i++ {
		fmt.Print(i*i, " ")
	}
}
