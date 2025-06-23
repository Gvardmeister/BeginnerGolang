package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Введите количество членов ряда: ")
	fmt.Scan(&n)

	sum := (1 - math.Pow(-0.5, float64(n))) / 1.5
	fmt.Printf("Сумма %d членов ряда: %.6f\n", n, sum)
}
