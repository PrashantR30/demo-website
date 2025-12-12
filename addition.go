package main

// AddThree adds three integers together.
func AddThree(a, b, c int) int {
	return a + b + c
}

// SumSlice returns the total sum of all integers in a slice.
// If the slice is empty, it returns 0.
func SumSlice(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

