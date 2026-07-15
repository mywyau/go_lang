package searching

// exporting and non-private via Capitalisation

func TwoSum(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		needed := target - num

		// this syntax combines the map look up with the if condition
		if previousIndex, found := seen[needed]; found {
			return []int{previousIndex, i}
		}

		seen[num] = i
	}

	return nil
}

func TwoSumDiffSyntax(nums []int, target int) []int {
	seen := map[int]int{}

	for i, num := range nums {
		needed := target - num

		previousIndex, found := seen[needed]
		if found {
			return []int{previousIndex, i}
		}

		seen[num] = i
	}

	return nil
}
