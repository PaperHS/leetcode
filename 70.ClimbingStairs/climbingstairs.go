package main

import "fmt"

func main() {

	//	fmt.Println("ret:", climbStairs(3))

	//	fmt.Println("ret:", climbStairs(5))
	fmt.Println("ret:", choose(23, 21))
	fmt.Println("ret:", climbStairs(44)) //
	//
}

func climbStairs2(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		return climbStairs(n-1) + climbStairs(n-2)
	}
}
func climbStairs(n int) int {
	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else {
		p1 := 1
		p2 := 2
		total := 0
		for i := 3; i <= n; i++ {
			total = p1 + p2
			p1 = p2
			p2 = total
		}
		return total
	}
}

func choose(total int, sec int) int {
	return 1
}

func powerf2(x int, n int) int {
	if n == 0 {
		return 1
	} else {
		return x * powerf2(x, n-1)
	}
}
