package main

import "fmt"

const dayYear = 365

func main() {
	var age int
	fmt.Scan(&age)

	fmt.Println(dayYear * age)
}
