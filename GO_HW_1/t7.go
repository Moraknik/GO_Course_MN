package main

import "fmt"

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		if isPrime(i) {
			fmt.Println(i)
		}
	}
}
