package warmup

func birthdayCakeCandles(candles []int32) int32 {
	tallest := candles[0]
	count := int32(0)

	for _, candle := range candles {
		if candle > tallest {
			tallest = candle
			count = 1
		} else if candle == tallest {
			count++
		}
	}

	return count
}
