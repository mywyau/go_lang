package warmup

func aVeryBigSum(ar []int64) int64 {

	var sum int64 = 0

	for _, num := range ar {
		sum += num
	}

	return sum
}
