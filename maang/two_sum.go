package maang

func twoSumMaang(nums []int, target int) []int {

	seen := map[int]int{}

	for i, num := range nums {
		needed := target - num
		if previousIndex, found := seen[needed]; found {
			return []int{previousIndex, i}
		}
		seen[num] = i
	}

	return nil
}
