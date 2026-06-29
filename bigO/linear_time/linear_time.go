package bigo

// O(n) - linear time
// This means the work grows in line with the input size.
// if nums has 5 items the the loop runs 5 times
// if it has 1000 items then it runs 1000 times.



func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}


// string linear time

func countLetters(s string) int {
	count := 0

	for range s {
		count++
	}

	return count
}