package main

import "testing"

func TestContains(t *testing.T) {
	nums := []int{0,12,22,23,33,46,50,59,61,65,71}
	if !contains(nums, 23) ||contains(nums,21) {
		t.Error("包含错误")
	}
}

func TestSum(t *testing.T) {
	nums := []int{0,12,22,23,33,46,50,59,61,65,66,67,71}
	if resultError(twoSum(nums,23),0,3) {
		t.Error("查找错误")
	}
	if resultError(twoSum(nums,33+59),4,7) {
		t.Error(twoSum(nums,33+59))
	}

	if resultError(twoSum(nums,61+59),8,7) {
		t.Error(twoSum(nums,61+59))
	}

	if resultError(twoSum(nums,23+50),3,6) {
		t.Error(twoSum(nums,23+50))
	}
	if resultError(twoSum(nums,23+50),3,6) {
		t.Error(twoSum(nums,23+50))
	}
	if resultError(twoSum(nums,23+65),3,9) {
		t.Error(twoSum(nums,23+50))
	}
	nums = []int{50,54,59,61,71,89,90,91,92}
	if resultError(twoSum(nums,50+90),0,6) {
		t.Error(twoSum(nums,50+90))
	}

	nums = []int{3,2,4}
	if resultError(twoSum(nums,6),1,2) {
		t.Error("错误")
	}

	nums = []int{3,2,2}
	if resultError(twoSum(nums,4),1,2) {
		t.Error("错误")
	}
}

func resultError(result []int, a int, b int) bool {
	return !contains(result,a) || !contains(result,b)
}
