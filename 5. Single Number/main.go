package main

func singleNumber(nums []int) int {
	keys := map[int]struct{}{}
	for _, num := range nums {
		if _, found := keys[num]; found {
			delete(keys, num)
		} else {
			keys[num] = struct{}{}
		}
	}
	for k := range keys {
		return k
	}
	return -1
}
