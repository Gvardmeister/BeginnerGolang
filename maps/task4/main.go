package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	words := strings.Fields(input)
	fmt.Println(takeMap(words))
}

func takeMap(words []string) map[string]int {
	newMap := make(map[string]int)

	for key, value := range words {
		if _, ok := newMap[value]; !ok {
			newMap[value] = key
		}
	}

	return newMap
}
