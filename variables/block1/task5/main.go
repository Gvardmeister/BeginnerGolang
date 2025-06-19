package main

import "fmt"

func main() {
	var min int
	fmt.Scan(&min)

	hours := min / 60
	minLost := min - (hours * 60)

	fmt.Printf("%d часа %d минут", hours, minLost)
}
