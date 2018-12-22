package main

import (
  "fmt"
)

func IsPalindrome(str string) bool {

  for i, j := 0, len(str); i<len(str) && i >= j;  i, j = i+1, j-1 {
    //if( str[i] != str[j] ) { return false }
    fmt.Printf("%c", str[i])
    fmt.Println("This is a test!")
  }
 
  return true
}

func main() {
  IsPalindrome("a")
  IsPalindrome("aba")
}
