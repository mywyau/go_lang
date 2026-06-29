package bigo

// O(n)
// This means the work grows in line with the input size.


func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}