package main

import "fmt"

func isVowel(ch string) bool {
	return (ch == "a") || (ch == "o") || (ch == "e") || (ch == "i") || (ch == "u") || (ch == "A") || (ch == "O") || (ch == "E") || (ch == "I") || (ch == "U")

}

func main() {
	var ch string
	fmt.Scanln(&ch)
	if isVowel(ch) {
		fmt.Println(ch, "is vowel")
	} else {
		fmt.Println(ch, "is consonant")
	}
}
