package main

import "fmt"

func main() {
	widths := []int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s := "bbbcccdddaaa"
	fmt.Println("ret:", numberOfLines(widths, s))
	widths = []int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}
	s = "abcdefghijklmnopqrstuvwxyz"
	fmt.Println("ret:", numberOfLines(widths, s))
}

func numberOfLines(widths []int, S string) []int {
	if len(S) == 0 {
		return []int{0, 0}
	}
	count := 0
	line := 1
	for _, a := range S {
		idx := int(a) - int('a')
		if count+widths[idx] > 100 {
			count = 0
			line++
		}
		count += widths[idx]
	}
	return []int{line, count}
}
