package main

func reverse(x int) int {
	res := 0
	isPos := 1
	if x < 0 {
		isPos = -1
		x = -1 * x
	}
	for {
		if x == 0 {
			break
		}
		res = res*10 + x%10
		if res > 2147483647 {
			return 0
		}
		x /= 10
	}
	return res * isPos

}