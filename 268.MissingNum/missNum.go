package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("out:", missingNumber(a))
	fmt.Println("out:", missingNumber([]int{0, 1, 2, 3, 4, 5, 6, 8, 9}))
	fmt.Println("out:", missingNumber([]int{0, 1, 2, 3, 4, 6, 7, 8, 9}))

}

func missingNumber(nums []int) int {
	max := len(nums)
	total := max * (max + 1) / 2
	for _, n := range nums {
		total -= n
	}
	return total
}
