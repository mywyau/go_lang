package searching

// This is LeetCode 35.
// Given a sorted array, return the index of the target. If it does not exist, return where it should be inserted.

func FindInsert(nums []int, target int) int {
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

	return left    //in binary search we return -1 but here we return left
}

// Why return left?
// Because when the loop finishes, left points to the correct insert position.
