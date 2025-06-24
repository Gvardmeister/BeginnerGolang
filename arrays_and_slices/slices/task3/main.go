package main

import "fmt"

func main() {
	sl := []int{10, 20, 30, 40}
	first := sl[0]

	for i := 0; i < len(sl)-1; i++ {
		sl[i] = sl[i+1]
	}
	sl[len(sl)-1] = first

	fmt.Println(sl)
}
