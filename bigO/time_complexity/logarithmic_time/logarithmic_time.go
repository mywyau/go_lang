package logarithmictime

// O(log n)
// This means the input gets smaller very quickly, usually by being cut in half each step.

// Each loop throws away half the remaining array.
// Example with 16 items:
// 16 -> 8 -> 4 -> 2 -> 1
// Only about 4 checks.

func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}