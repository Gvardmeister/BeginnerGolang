package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 4, 5}
	var n int = 2

	resultSl := append(sl[:n], sl[n+1:]...)
	fmt.Println(resultSl)
}
