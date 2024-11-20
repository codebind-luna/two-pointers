package main

func sortArrayByParityII(nums []int) []int {
	even, odd := 0, 1
	for even < len(nums) && odd < len(nums) {
		if nums[even]%2 != 0 {
			nums[even], nums[odd] = nums[odd], nums[even]
			odd += 2
		} else {
			even += 2
		}
	}
	return nums
}
