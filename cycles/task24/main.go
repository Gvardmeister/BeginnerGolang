package main

import "fmt"

func main() {
	const (
		pricePen    = 10
		pricePencil = 5
		priceEraser = 2
		totalItems  = 30
		totalSum    = 100
	)

	fmt.Println("Возможные варианты (ручки, карандаши, ластики):")
	for a := 0; a <= totalItems; a++ {
		for b := 0; b <= totalItems-a; b++ {
			c := totalItems - a - b
			if pricePen*a+pricePencil*b+priceEraser*c == totalSum {
				fmt.Printf("Ручки: %d, Карандаши: %d, Ластики: %d\n", a, b, c)
			}
		}
	}
}
