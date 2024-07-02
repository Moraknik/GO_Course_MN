package main

import "fmt"

func add(val1 int, val2 int) int {
	return val1 + val2
}

func main() {
	var val1, val2 int
	fmt.Scanln(&val1)
	fmt.Scanln(&val2)
	fmt.Println(add(val1, val2))
}
