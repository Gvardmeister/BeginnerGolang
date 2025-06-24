package main

import "fmt"

func main() {
	sl := []int{1, 2, 3, 2, 1}
	sl2 := []int{1, 2, 3, 3, 1}

	fmt.Println(isPalindrom(sl))
	fmt.Println(isPalindrom(sl2))
}

func isPalindrom(sl []int) bool {
	for i := 0; i < len(sl)/2; i++ {
		if sl[i] != sl[len(sl)-1-i] {
			return false
		}
	}

	return true
}
