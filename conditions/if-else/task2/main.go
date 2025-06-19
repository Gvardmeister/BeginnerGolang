package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)

	if age >= 18 {
		fmt.Println("Доступ разрешен")
	} else {
		fmt.Println("Доступ запрещен")
	}
}
