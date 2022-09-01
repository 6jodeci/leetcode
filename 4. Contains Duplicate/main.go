package main

func containsDuplicate(nums []int) bool {
	keys := map[int]bool{}
	for i := range nums {
		if keys[nums[i]] {
			return true
		}
		keys[nums[i]] = true
	}
	return false
}
