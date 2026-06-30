package logarithmictime

import "sort"

// O(n log n)
// This often appears in efficient sorting algorithms.

// The average/worst complexity of Go’s standard sorting algorithms is usually discussed as around:
// O(n log n)

func sorting_example(nums []int) {
	sort.Ints(nums)
}


