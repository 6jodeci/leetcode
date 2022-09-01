package main

func twoSum(nums []int, target int) []int {
	keys := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, ok := keys[target-nums[i]]
		if ok {
			result := []int{j, i}
			return result
		}
		keys[nums[i]] = i
	}
	result := []int{-1, -1}
	return result
}
