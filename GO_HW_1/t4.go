package main

import "fmt"

func main() {
	var val1, val2, val3 int
	fmt.Scanln(&val1)
	fmt.Scanln(&val2)
	fmt.Scanln(&val3)
	fmt.Println(max(val1, val2, val3))
}
