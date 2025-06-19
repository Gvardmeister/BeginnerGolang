package main

import "fmt"

func main() {
	fmt.Println(greet("Alex"))
}

func greet(name string) string {
	return "Привет, " + name + "!"
}
