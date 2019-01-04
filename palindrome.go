package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(str string) bool {

	for i, j := 0, len(str)-1; i < len(str); i, j = i+1, j-1 {
		if unicode.ToLower(rune(str[i])) != unicode.ToLower(rune(str[j])) {
			fmt.Println(str, "False")
			return false
		}
	}
	fmt.Println(str, "True")
	return true
}
