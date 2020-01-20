package arrays_and_slices

// Sum sums a variable number of arguments given as input
func Sum(nums []int) int {
	var total int = 0
	for _, n := range nums {
		total += n
	}
	return total
}
