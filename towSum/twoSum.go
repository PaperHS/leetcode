package main

import (
	"sort"
	"math/rand"
)

func main() {
	var nums []int
	for true {
		x := rand.Intn(100)
		if !contains(nums, x) {
			nums = append(nums, x)
		}
		if len(nums) > 20 {
			break
		}
	}
	twoSum(nums, 78)
}

func twoSum(nums []int, sum int) []int {
	ori := make([]int, len(nums))
	copy(ori, nums)
	sort.Ints(nums)
	left := index(nums, sum/2)
	right := left + 1
	for nums[left]+nums[right] != sum {
		if nums[left]+nums[right] > sum {
			left--
			continue
		}
		if nums[left]+nums[right] < sum {
			right++
			continue
		}
	}
	return pos2(ori, nums[left], nums[right])
}

func pos2(nums []int, x int, y int) []int {
	ret := []int{-1, -1}
	for i, e := range nums {
		if e == x && ret[0] == -1 {
			ret[0] = i
		} else if e == y && ret[1] == -1 {
			ret[1] = i
		}
		if ret[0] != -1 && ret[1] != -1 {
			return ret
		}
	}
	return ret
}

func index(nums []int, x int) int {
	if x < nums[0] {
		return 0
	}
	if x > nums[len(nums)-1] {
		return len(nums) - 1
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == x || (nums[i] < x && nums[i+1] > x) {
			return i
		}
	}
	return -1
}

//func sum(nums []int ,sum int) []int{
//	ori := make([]int,len(nums))
//	copy(ori,nums)
//	fmt.Println(ori)
//	sort.Ints(nums)
//	fmt.Println(nums)
//	left := index(nums, sum/2)
//	right := left +1
//	l := len(nums)
//	for i := left;i >= 0 ; i-- {
//		for j := right;  j < l; j++  {
//			if sum == nums[i] + nums[j]{
//				fmt.Println("i,j=",i,j,nums[i],nums[j],index(ori,nums[i]))
//				return pos(ori,nums[i],nums[j])
//			}
//		}
//	}
//	//for nums[left] + nums[right] != sum{
//	//	if nums[left] + nums[right] > sum {
//	//		left--
//	//		continue
//	//	}
//	//	if nums[left] + nums[right] < sum {
//	//		right++
//	//		continue
//	//	}
//	//	}
//
//	return []int{left,right}
//}
//
//func index(nums []int ,x int) int {
//	if x < nums[0] {
//		return 0
//	}
//	if x > nums[len(nums)-1] {
//		return len(nums)-1
//	}
//	for i := 0; i < len(nums) -1 ; i++ {
//		if nums[i] == x || (nums[i] < x && nums[i+1] > x)  {
//			return i
//		}
//	}
//	return -1
//}
//
//func pos(nums []int, x int, y int) []int {
//	ret := []int{-1,-1}
//	for i,e := range nums {
//		if e == x && ret[0] == -1{
//			ret[0] = i
//		}else if e == y && ret[1] == -1 {
//			ret[1] = i
//		}
//		if ret[0] != -1 && ret[1] != -1 {
//			return ret
//		}
//	}
//	return ret
//}

func contains(nums []int, x int) bool {
	for _, e := range nums {
		if e == x {
			return true
		}
	}
	return false
}
