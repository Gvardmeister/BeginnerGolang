package main

import (
	"fmt"
	"strings"
)

func main() {
	sl := []string{"Go", "is", "awesome"}
	result := strings.Join(sl, " ")
	fmt.Println(result)
}
