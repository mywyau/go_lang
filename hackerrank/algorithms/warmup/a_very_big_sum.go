package warmup

// we first define a sum as int64
// then we iterate as normal thorugh the array, adding to the sum 
// finally we return the sum 
// done

func aVeryBigSum(ar []int64) int64 {

	var sum int64 = 0

	for _, num := range ar {
		sum += num
	}

	return sum
}


