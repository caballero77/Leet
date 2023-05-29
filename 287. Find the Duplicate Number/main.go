package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}
	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return fast
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
}
