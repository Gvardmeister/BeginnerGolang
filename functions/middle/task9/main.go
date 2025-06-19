package main

import "fmt"

var counter int

func increment() {
	counter++
}

func main() {
	increment()
	increment()
	increment()
	fmt.Println("Значение counter:", counter)
}
