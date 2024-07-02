package main

import "fmt"

func stringReverse(st string) string {
	return st
}

func main() {
	var st string
	fmt.Scanln(&st)
	fmt.Println(stringReverse(st))
}
