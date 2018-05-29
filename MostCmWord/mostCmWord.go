package main

import (
	"fmt"
	"strings"
)

func main() {
	paragraph := "Bob hit a ball, the hit BALL flew far after it was hit."
	banned := []string{"hit"}
	fmt.Println("ret:", mostCommonWord(paragraph, banned))
	fmt.Println("ret:", mostCommonWord("A.", []string{""}))
}

func mostCommonWord(paragraph string, banned []string) string {
	s := strings.ToLower(paragraph)
	s = strings.Replace(s, ".", "", -1)
	s = strings.Replace(s, ";", "", -1)
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, "'", "", -1)
	s = strings.Replace(s, "?", "", -1)
	s = strings.Replace(s, "!", "", -1)
	if !strings.ContainsRune(s, ' ') {
		return s
	}
	dmap := make(map[string]int)
	for _, word := range strings.Split(s, " ") {
		dmap[word] += 1
	}
	for _, v := range banned {
		dmap[v] = 0
	}
	max := 0
	value := ""
	for k, v := range dmap {
		if v > max {
			max = v
			value = k
		}
	}
	return value
}
