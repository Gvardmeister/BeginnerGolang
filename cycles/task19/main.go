package main

import "fmt"

func main() {
	fmt.Println("Совершенные числа от 1 до 10000:")

	for n := 1; n <= 10000; n++ {
		sum := 0
		for d := 1; d < n; d++ {
			if n%d == 0 {
				sum += d
			}
		}
		if sum == n {
			fmt.Println(n)
		}
	}
}
