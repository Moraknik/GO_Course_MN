package main

import "fmt"

func main() {
	var val int
	fmt.Scanln(&val)
	if val%2 == 0 {
		fmt.Println(val, "is even")
	} else {
		fmt.Println(val, " is odd")
	}
}
