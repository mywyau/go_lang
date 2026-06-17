package algorithms

import "fmt"

/*
	Mini-Max Sum

	Given exactly 5 integers, calculate:
	- the minimum sum of 4 numbers
	- the maximum sum of 4 numbers

	Instead of generating every possible 4-number sum, use this idea:

	    totalSum = sum of all 5 numbers

	    minimum 4-number sum = totalSum - largest number
	    maximum 4-number sum = totalSum - smallest number

	Example:

	    arr = []int32{1, 2, 3, 4, 5}

	    totalSum = 15
	    smallest = 1
	    largest = 5

	    minSum = 15 - 5 = 10
	    maxSum = 15 - 1 = 14

	Output:

	    10 14

	Use int64 for the sum because the final result can be larger
	than the range of int32.
*/

// hackerrank solution

// func miniMaxSum(arr []int32) {
// 	var totalSum int64 = 0

// 	smallest := int64(arr[0])
// 	largest := int64(arr[0])

// 	for _, num := range arr {
// 		value := int64(num)

// 		totalSum += value

// 		if value < smallest {
// 			smallest = value
// 		}

// 		if value > largest {
// 			largest = value
// 		}
// 	}

// 	minSum := totalSum - largest
// 	maxSum := totalSum - smallest

// 	fmt.Println(minSum, maxSum)
// }

// testable version

// func calculateMiniMaxSum(arr []int32) (int64, int64) {
// 	var totalSum int64 = 0

// 	smallest := int64(arr[0])
// 	largest := int64(arr[0])

// 	for _, num := range arr {
// 		value := int64(num)

// 		totalSum += value

// 		if value < smallest {
// 			smallest = value
// 		}

// 		if value > largest {
// 			largest = value
// 		}
// 	}

// 	minSum := totalSum - largest
// 	maxSum := totalSum - smallest

// 	return minSum, maxSum
// }

func calculateMiniMaxSum(arr []int32) (int64, int64) {
	// Start with the first value as our current smallest and largest.
	// We use int64 because the sum may be larger than int32 can safely hold.
	var totalSum int64 = 0
	smallest := int64(arr[0])
	largest := int64(arr[0])

	// Loop through every number.
	// Add it to the total sum and update smallest/largest when needed.
	for _, num := range arr {
		value := int64(num)
		totalSum += value

		if value < smallest {
			smallest = value
		}

		if value > largest {
			largest = value
		}
	}

	// Minimum 4-number sum: remove the largest value.
	// Maximum 4-number sum: remove the smallest value.
	minSum := totalSum - largest
	maxSum := totalSum - smallest

	return minSum, maxSum
}

func miniMaxSum(arr []int32) {
	minSum, maxSum := calculateMiniMaxSum(arr)
	fmt.Println(minSum, maxSum)
}
