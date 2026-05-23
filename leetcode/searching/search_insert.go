package searching

// This is LeetCode 35.
// Given a sorted array, return the index of the target. If it does not exist, return where it should be inserted.

func SearchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

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

	return left
}

// Why return left?
// Because when the loop finishes, left points to the correct insert position.
