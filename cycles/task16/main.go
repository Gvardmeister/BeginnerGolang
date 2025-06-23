package main

import "fmt"

func main() {
	var n, d int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)
	fmt.Print("Введите цифру для удаления: ")
	fmt.Scan(&d)

	result := 0
	multiplier := 1

	for n > 0 {
		digit := n % 10
		if digit != d {
			result += digit * multiplier
			multiplier *= 10
		}
		n /= 10
	}

	fmt.Println("Результат:", result)
}
