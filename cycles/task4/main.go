package main

import (
	"fmt"
)

func main() {
	start := -5.0
	end := 5.0

	for start <= end {
		y := 5 - start*start/2
		fmt.Printf("\nx = %.1f, y = %.2f", start, y)
		start += 0.5
	}
}
