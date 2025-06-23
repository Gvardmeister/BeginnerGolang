package main

import "fmt"

func main() {
	var limit, targetSum int
	fmt.Print("Введите предел трех чисел: ")
	fmt.Scan(&limit)
	fmt.Print("Введите число-сумму: ")
	fmt.Scan(&targetSum)

	fmt.Println("Комбинации трех чисел, сумма которых равна", targetSum, ":")

	for a := 0; a <= limit; a++ {
		for b := 0; b <= limit; b++ {
			for c := 0; c <= limit; c++ {
				if a+b+c == targetSum {
					fmt.Printf("(%d, %d, %d)\n", a, b, c)
				}
			}
		}
	}
}
