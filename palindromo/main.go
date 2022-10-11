package main

import (
	"fmt"
	"strings"
)

func isPalindromo(word string) {
	toLowerWord := strings.ToLower(word)
	wordLength := len(toLowerWord)
	isPalindrome := true
	for i, j := wordLength-1, 0; i >= 0; i, j = i-1, j+1 {
		if toLowerWord[i] != toLowerWord[j] {
			isPalindrome = false
		}
	}
	if isPalindrome {
		fmt.Printf("The word: '%v': is palindrome", word)
	} else {
		fmt.Printf("The word: '%v': is not palindrome", word)
	}
}

func main() {
	word := "Am"
	isPalindromo(word)
}
