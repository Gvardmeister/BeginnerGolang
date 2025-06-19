package main

import (
	"fmt"
	"math"
)

func main() {
	var numA, numB float64
	fmt.Scan(&numA, &numB)

	fmt.Println(math.Abs(numA) - math.Abs(numB))
}
