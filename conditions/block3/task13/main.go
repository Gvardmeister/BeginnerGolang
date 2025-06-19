package main

import "fmt"

func main() {
	var numA, numB int
	fmt.Scan(&numA, &numB)

	if numA > numB {
		fmt.Println("Первое больше")
	} else {
		fmt.Println("Второе больше")
	}
}
