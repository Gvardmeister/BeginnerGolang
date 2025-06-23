package main

import (
	"fmt"
)

func main() {
	for {
		var op string
		fmt.Print("Введите операцию (+, -, *, /) или 0 для выхода: ")
		fmt.Scan(&op)

		if op == "0" {
			fmt.Println("Программа завершена.")
			break
		}

		if op != "+" && op != "-" && op != "*" && op != "/" {
			fmt.Println("Некорректный знак операции. Попробуйте снова.")
			continue
		}

		var x, y float64
		fmt.Print("Введите два числа (x y): ")
		_, err := fmt.Scan(&x, &y)
		if err != nil {
			fmt.Println("Ошибка ввода чисел. Попробуйте снова.")
			continue
		}

		var result float64
		var valid bool = true

		switch op {
		case "+":
			result = x + y
		case "-":
			result = x - y
		case "*":
			result = x * y
		case "/":
			if y == 0 {
				fmt.Println("Ошибка: деление на ноль!")
				valid = false
			} else {
				result = x / y
			}
		}

		if valid {
			fmt.Printf("Результат: %.4f\n", result)
		}
	}
}
