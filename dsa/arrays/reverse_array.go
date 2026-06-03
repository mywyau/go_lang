package arrays

func reverseArray(a []int32) []int32 {

	var reversed []int32

    for i := len(a) - 1; i >= 0; i-- {
		reversed = append(reversed, a[i])
	}

	return reversed
}
