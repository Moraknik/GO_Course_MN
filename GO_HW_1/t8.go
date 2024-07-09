package main

import "fmt"

func stringReverse(st string) string {
	stArr := []rune(st)
	for i, j := 0, len(stArr)-1; i < j; i, j = i+1, j-1 {
		stArr[i], stArr[j] = stArr[j], stArr[i]
	}
	return string(stArr)
}

func main() {
	var st string
	fmt.Scanln(&st)
	fmt.Println(stringReverse(st))
}
