package main
package primes

import (
	"fmt"
	"primes"
)

func main() {
	var n int
	fmt.Scanln(&n)
	for i = 1; i <= n; ++i {
		if IsPrime(i) {
			fmt.Print(i)
			}
		}
	}
}
