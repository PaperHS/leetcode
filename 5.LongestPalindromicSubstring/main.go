package main

import (
	"fmt"
)

func main() {

	fmt.Println("l::::%", longestPalindrome("abbcccccbb"))
	fmt.Println("l::::%", longestPalindrome("qwerabbxcxcbb"))
	fmt.Println("l::::%", longestPalindrome("abasfdbccbb"))
	fmt.Println("l::::%", longestPalindrome("abbccsafbb"))
	fmt.Println("l::::%", longestPalindrome("abbccafdbb"))

}

func lp2(s string) string {

}

func longestPalindrome(s string) string {
	left := 0
	right := len(s) - 1
	addleft := true
	var rst []byte
	singleLast := false
	for {
		if left > right {
			break
		}
		if left == right {
			rst = append(rst, s[left])
			singleLast = true
			break
		} else if s[left] == s[right] {
			rst = append(rst, s[left])
			left++
			right--
		} else {
			if addleft {
				left++
			} else {
				right--
			}
			addleft = !addleft
			rst = rst[:0]
		}
	}

	if singleLast {
		rv := reverse(rst)[1:]
		return string(append(rst, rv...))
	} else {
		return string(append(rst, reverse(rst)...))
	}
}

func reverse(sl []byte) []byte {
	ret := make([]byte, len(sl))
	for i, c := range sl {
		ret[len(sl)-i-1] = c
	}
	return ret
}
