package algorithms

import "fmt"

/*
	Plus Minus

	Given an array of integers, calculate the ratio of:
	- positive numbers
	- negative numbers
	- zero values

	Ratio means:

	    count of category / total number of items

	Example:

	    arr := []int32{-4, 3, -9, 0, 4, 1}

	    positives: 3 values -> 3, 4, 1
	    negatives: 2 values -> -4, -9
	    zeroes:    1 value  -> 0
	    total:     6 values

	    positive ratio = 3 / 6 = 0.500000
	    negative ratio = 2 / 6 = 0.333333
	    zero ratio     = 1 / 6 = 0.166667

	HackerRank expects each ratio printed on a new line,
	with exactly 6 decimal places.
*/

func plusMinus(arr []int32) {
	var positiveCount int32 = 0
	var negativeCount int32 = 0
	var zeroCount int32 = 0

	total := int32(len(arr))

	for _, num := range arr {
		if num > 0 {
			positiveCount++
		} else if num < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	fmt.Printf("%.6f\n", float64(positiveCount)/float64(total))
	fmt.Printf("%.6f\n", float64(negativeCount)/float64(total))
	fmt.Printf("%.6f\n", float64(zeroCount)/float64(total))
}

// testable version

func calculateRatios(arr []int32) (float64, float64, float64) {
	var positiveCount int32 = 0
	var negativeCount int32 = 0
	var zeroCount int32 = 0

	total := int32(len(arr))

	for _, num := range arr {
		if num > 0 {
			positiveCount++
		} else if num < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
	}

	positiveRatio := float64(positiveCount) / float64(total)
	negativeRatio := float64(negativeCount) / float64(total)
	zeroRatio := float64(zeroCount) / float64(total)

	return positiveRatio, negativeRatio, zeroRatio
}
