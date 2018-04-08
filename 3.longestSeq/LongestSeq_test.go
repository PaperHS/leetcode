package logist

import (
	"fmt"
	"testing"
)

func Test_longestSeq(t *testing.T) {
	fmt.Printf("start \n")
	var seq0 string
	seq0 = "abcdefghijklmnopqrstuvwxyz"
	testSeq(seq0, 26, t)
}

func Test_132(t *testing.T) {
	// 1 3 2
	testSeq("abcadefghijklmnopqarstuvwxyz", 24, t)
}
func Test_123(t *testing.T) {
	// 1 3 2
	testSeq("cadefghijklmnopqarstuvwxyz", 24, t)
}
func Test_abba(t *testing.T) {
	// a bb a
	testSeq("abcdefghijklmnopqbrstuvbwxyza", 21, t)
	testSeq("acdefghijklmnopqbrstuvbwxyza", 22, t)
}
func Test_abab(t *testing.T) {
	// a b a   b
	testSeq("abcdefghaijklmnopqrstuvdwxyz", 24, t)
}
func Test_333(t *testing.T) {
	// 1 3 2
	testSeq("abcadefghijklmnopqarstuvwxyz", 24, t)
	testSeq("", 0, t)
	testSeq("a", 1, t)
	testSeq("ab", 2, t)
	testSeq("aba", 2, t)

}

func testSeq(seq0 string, l int, t *testing.T) {
	length := lengthOfLongestSubstring(seq0)
	if length != l {
		t.Errorf("seq max len:%d", length)
	} else {
		fmt.Println("==================")
	}
}
