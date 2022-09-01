package main

func plusOne(digits []int) []int {
    carry := 1

    for i:= len(digits) - 1; i >= 0; i-- {
        digits[i] += carry

        if digits[i] < 10 {
            return digits
        } else {
            digits[i] = 0
            carry = 1
        }
    }

    if carry != 0 {
        return append([]int{1}, digits...)
    } else {
        return digits
    }
}