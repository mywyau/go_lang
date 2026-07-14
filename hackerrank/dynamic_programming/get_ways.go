package dynamicprogramming


// coin change problem

func getWays(amount int32, coins []int32) int64 {
	
	ways := make([]int64, amount+1)

	// There is one way to make 0:
	// choose no coins.
	ways[0] = 1

	for _, coin := range coins {
		for currentAmount := coin; currentAmount <= amount; currentAmount++ {
			ways[currentAmount] += ways[currentAmount-coin]
		}
	}

	return ways[amount]
}


