package main

func rotate(nums []int, k int) {
	k = k % len(nums)
	result := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = result[i]
	}
}
