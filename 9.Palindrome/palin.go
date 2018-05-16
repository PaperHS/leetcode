package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(isPalindrome(123123))
	fmt.Println(isPalindrome(12321))
	fmt.Println(isPalindrome(-123321))
	fmt.Println(isPalindrome(1234566654321))
	fmt.Println(isPalindrome(123123))
	fmt.Println(isPalindrome(123123))
	fmt.Println(isPalindrome(123123))

}

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	right := len(s) - 1
	for i := 0; i <= right/2; i++ {
		if s[i] != s[right-i] {
			return false
		}
	}
	return true
}
