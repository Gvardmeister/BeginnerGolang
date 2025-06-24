package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5} // len(5), cap(5)
	t := s[1:4]               // len(3), cap(4)
	t[0] = 99

	fmt.Println(s) // {1, 99, 3, 4, 5}
	fmt.Println(t) // {99, 3, 4}
}
