package main

import "fmt"

func main() {
	var a, b, c int
	var d int

	fmt.Print("Введите два натуральных числа: ")
	fmt.Scan(&a, &b)

	for a > 0 {
		d = a % 10
		a = a / 10
		c = b

		for c > 0 {
			if c%10 == d {
				fmt.Printf("%d ", d)
				break
			}
			c = c / 10
		}
	}
	fmt.Println()
}
