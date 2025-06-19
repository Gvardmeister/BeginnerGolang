package main

import "fmt"

func main() {
	fmt.Println(length("hello"))
}

func length(s string) int {
	return len(s)
}
