package arrays_and_slices

// Sum sums a slice of ints, returning a single int
func Sum(nums []int) int {
	var total int = 0
	for _, n := range nums {
		total += n
	}
	return total
}

// SumAll computes the sum of each slice passed in as input
// returning the result in a slice of ints where the index
// corresponds to the positional ordering of the int slice
// passed as input
func SumAll(nums ...[]int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = Sum(n)
	}
	return result
}
