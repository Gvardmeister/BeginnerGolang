package main

import (
	"fmt"
	"math"
)

func main() {
	var number, step int
	fmt.Scan(&number, &step)

	num := math.Pow(float64(number), float64(step))
	fmt.Println(num)
}
