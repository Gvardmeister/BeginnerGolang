package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const min, max = 1, 100

	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(max-min+1) + min

	var b int
	fmt.Printf("Угадайте число от %d до %d\n", min, max)

	for {
		fmt.Print("Введите число: ")
		_, err := fmt.Scan(&b)
		if err != nil {
			fmt.Println("Ошибка ввода. Попробуйте снова.")
			continue
		}

		if b > a {
			fmt.Println("Много")
		} else if b < a {
			fmt.Println("Мало")
		} else {
			fmt.Println("Угадал!")
			break
		}
	}
}
