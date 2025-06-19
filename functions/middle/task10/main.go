package main

import "fmt"

func main() {
	fmt.Println(scopedVar())
}

func scopedVar() int {
	x := 5

	return x
}
