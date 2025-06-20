package main

import "fmt"

// Гипотеза Сиракуз
func main() {
	for i := 20; i <= 30; i++ {
		n := i
		fmt.Printf("Путь для %d ", n)
		for n != 1 {
			if n%2 == 0 {
				n = n / 2
			} else {
				n = (3*n + 1) / 2
			}
			fmt.Printf("%d ", n)
		}
		fmt.Println()
	}
}
