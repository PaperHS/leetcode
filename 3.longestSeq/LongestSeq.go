package logist

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	var slice1 string
	l := 0
	//	lastPos := 0
	for i, c := range s {
		//		fmt.Printf("i:%d c:%d \n", i, c)
		index := index(slice1, c)
		if index >= 0 || i == len(s)-1 {
			nl := 0
			if i == len(s)-1 && index < 0 {
				nl = len(slice1) + 1
			} else {
				nl = len(slice1)
			}
			if nl > l {
				l = nl
			}
			fmt.Printf("slice=%s \n", slice1)
			slice1 = slice1[index+1:]

			//			fmt.Printf("pos : %d \n", i)
		}
		slice1 = slice1 + string(c)

	}
	return l
}
func index(s string, c rune) int {
	//	fmt.Println("s :" + s + "  c=" + string(c) + "\n")
	for i, m := range s {

		if m == c {
			//			fmt.Printf("m = " + string(m) + "\n")
			return i
		}
	}
	return -1
}
