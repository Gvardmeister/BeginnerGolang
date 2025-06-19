package main

import "fmt"

var name string = "Alex"

func main() {
	changeName()

	fmt.Println(name)
}

func changeName() {
	name = "Nasty"
}
