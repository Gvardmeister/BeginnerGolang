package main

import "fmt"

func main() {
	var numA, numB int
	fmt.Scan(&numA, &numB)

	if numA > numB {
		fmt.Println("Число А больше")
	} else if numA == numB {
		fmt.Println("Числа равны")
	} else {
		fmt.Println("Число B больше")
	}
}
