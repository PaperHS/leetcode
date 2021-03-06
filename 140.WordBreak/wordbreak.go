package main

import (
	"fmt"
	"strings"
)

func main() {
	//	s := "catsanddog"
	//	wordDict := []string{"cats", "dog", "sand", "and", "cat"}
	//	fmt.Println("ret:", wordBreak(s, wordDict))
	//	s = "pineapplepenapple"
	//	wordDict = []string{"apple", "pen", "applepen", "pine", "pineapple"}
	//	fmt.Println("ret:", wordBreak(s, wordDict))
	s := "aaaaaaa"
	wordDict := []string{"aaaa", "aa", "a"}
	out := wordBreak(s, wordDict)
	for _, x := range out {
		fmt.Println("out:", x)
	}

	wrong := []string{"aaaa aa a", "aa aaaa a", "a aaaa a", "aaaa a a", "aa aa a a", "aa a a", "a aa a", "a a a a", "a aaaa a a", "aaaa a a a", "aa aa a aa", "aa a a aa", "aa a a aa a", "a aa a aa", "a aa a aa a", "a a a aa a", "a a a aaaa", "aa aa a a a", "aa a a a a", "aa a a a a", "a aa a a a", "a aa a a a", "a a a a a", "a a a aa a", "a a a a a a", "aa a a a a a", "a aa a a a a", "a a a a a a", "a a a aa a a", "a a a a a a", "a a a a a a a"}
	right := []string{
		"a a a a a a a", "aa a a a a a", "a aa a a a a", "a a aa a a a", "aa aa a a a", "aaaa a a a", "a a a aa a a", "aa a aa a a", "a aa aa a a", "a aaaa a a", "a a a a aa a", "aa a a aa a", "a aa a aa a", "a a aa aa a", "aa aa aa a", "aaaa aa a", "a a aaaa a", "aa aaaa a", "a a a a a aa", "aa a a a aa", "a aa a a aa", "a a aa a aa", "aa aa a aa", "aaaa a aa", "a a a aa aa", "aa a aa aa", "a aa aa aa", "a aaaa aa", "a a a aaaa", "aa a aaaa", "a aa aaaa"}
	for _, a := range right {
		for _, b := range wrong {
			if a == b {
				break
			}
			//			if i == len(wrong)-1 {
			//				fmt.Println("a:", a)
			//				fmt.Println("b:", b)

			//		}
		}
	}
}

func doDivide(s string, wordDict []string, ret []string) {
	x := matchFirst(s, wordDict)
	for _, div = range x {

	}
}

func wordBreak(s string, wordDict []string) []string {
	ret := [][]string{{""}}
	remind := []int{0}
	remindbak := []string{}
	length := len(s)
	for !isAllEnd(remind, length) {
		clone := make([]int, len(remind))
		copy(clone, remind)
		for idx, subs := range clone {
			if subs < 0 || subs >= length {
				continue
			}
			x := matchFirst(subString(subs, s), wordDict)

			if len(x) > 0 {
				tmp := make([]string, len(ret[idx]))
				copy(tmp, ret[idx])
				tmpRe := remind[idx]
				for i, div := range x {
					if i > 0 {
						ret = append(ret, append(tmp, div))
						remind = append(remind, tmpRe+len(div))
					} else {
						ret[idx] = append(ret[idx], div)
						remind[idx] += len(div)
					}
				}
			} else {
				remind[idx] = -1
				break
			}
		}
	}
	for i, x := range ret {
		if remind[i] != -1 {
			remindbak = append(remindbak, strings.TrimSpace(strings.Join(x, " ")))
		}
	}
	return remindbak
}

func subString(index int, s string) string {

	sl := []rune(s)
	return string(sl[index:])
}

func isAllEnd(a []int, end int) bool {
	for _, x := range a {
		if x < end && x != -1 {
			return false
		}
	}
	return true
}

func matchFirst(s string, wordDict []string) []string {
	ret := []string{}
	for _, w := range wordDict {
		if strings.HasPrefix(s, w) {
			ret = append(ret, w)
		}
	}
	return ret
}
