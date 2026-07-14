package dynamicprogramming

func equal(arr []int32) int32 {
	minValue := arr[0]

	for _, value := range arr {
		if value < minValue {
			minValue = value
		}
	}

	best := int32(1 << 30)

	for target := minValue; target >= minValue-4; target-- {
		var operations int32 = 0

		for _, value := range arr {
			diff := value - target

			operations += diff / 5
			diff %= 5

			operations += diff / 2
			diff %= 2

			operations += diff
		}

		if operations < best {
			best = operations
		}
	}

	return best
}