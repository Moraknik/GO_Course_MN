package main

import "fmt"

func main() {
	var n int
	var factorial = 1
	fmt.Scanln(&n)
	for ; n != 0; n-- {
		factorial *= n

	}
	fmt.Println(factorial)
}
