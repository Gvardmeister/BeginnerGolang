package main

import (
	"fmt"
)

func main() {
	var n int
	count := 0

	for {
		fmt.Print("Введите число: ")
		fmt.Scan(&n)

		if n < 2 {
			break
		}

		if isPrime(n) {
			count++
		}
	}

	fmt.Println("Количество простых чисел:", count)
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
