package main

import "fmt"

func main() {
	fmt.Printf("%b\n", 0xffffffff)
	fmt.Printf("%b\n", 0x7fffffff)
	fmt.Printf("%b\n", -2147483647)
	fmt.Printf("%d\n", abs(-123))
	fmt.Printf("%d\n", abs(-2147483647))
	fmt.Println("hello%", abs(1234))
	fmt.Println("hello%", abs(-2^32))
	fmt.Println("hello%", abs(-23412312))
	fmt.Println("hello%", reverse(123))
	fmt.Println("hello%", reverse(-1234))
	fmt.Println("hello%", reverse(1))
	fmt.Println("hello%", reverse(10000))
	fmt.Println("hello%", reverse(100300))
	fmt.Println("hello%", reverse(2^31))
	fmt.Println("hello%", reverse(-2147483647))
}

func reverse(x int32) int32 {
	nege := x < 0
	_, ret := foo(abs(x), 0)
	if ret < 0 {
		return 0
	} else {
		if nege {
			return -1 * ret
		} else {
			return ret
		}
	}
}

func foo(in int32, out int32) (ii int32, oo int32) {
	if in == 0 {
		return 0, out
	} else {
		out = out*10 + in%10
		in = in / 10
		return foo(in, out)
	}
}

func abs(dd int32) int32 {
	y := dd >> 31
	return (dd ^ y) - y
}
