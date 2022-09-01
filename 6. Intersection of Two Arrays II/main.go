package main

func intersect(nums1 []int, nums2 []int) []int {
	keys := make(map[int]int)
	var arr []int
	for _, v := range nums1 {
		keys[v]++
	}

	for _, v := range nums2 {
		times, ok := keys[v]
		if ok && times > 0 {
			arr = append(arr, v)
			keys[v]--
		}
	}
	return arr
}
