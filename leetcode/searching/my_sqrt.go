package searching

func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid

		if square == x {
			return mid
		}

		if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}