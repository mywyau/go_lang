package warmup


// Find the maximum value in an array, then count how many times it appears. 

func birthdayCakeCandles(candles []int32) int32 {

	tallest := candles[0]   // initialise the tallest candle to be the initial value in the array as a guess
	count := int32(0)		// set the count for the number of tallest candles to be 0 at the beginning

	for _, candle := range candles {
		if candle > tallest {		// if the candle is taller than the current percieved tallest we need to update tallest and reset the count to one   
			tallest = candle
			count = 1
		} else if candle == tallest {   // we add the tallest candle to the count
			count++
		}
	}

	return count
}
