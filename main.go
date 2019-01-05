package main

import (
	"fmt"
)

func main() {

	fmt.Println("Checking Palindromes!  ******")

	isPalindrome("a")
	isPalindrome("aba")
	isPalindrome("abba")
	isPalindrome("Abba")
	isPalindrome("nbba")

	arr := []int{8, 3, 6, 5, 9, 1, 4, 0, 7, 2}

	fmt.Println("Before:", arr)
	bubbleSort(arr)
	fmt.Println("After:", arr)

	return
}
