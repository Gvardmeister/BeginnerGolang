package main

import "fmt"

func main() {
	var m, n int
	fmt.Print("Введите значения M и N: ")
	fmt.Scan(&m, &n)

	for i := m; i <= n; i++ {
		fmt.Printf("Делители числа %d: ", i)
		hasDivisors := false
		for d := 2; d <= i/2; d++ {
			if i%d == 0 {
				fmt.Printf("%d ", d)
				hasDivisors = true
			}
		}
		if !hasDivisors {
			fmt.Print("нет")
		}
		fmt.Println()
	}
}
