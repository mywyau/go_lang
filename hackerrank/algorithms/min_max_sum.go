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

	// we initiate the starting values
	var totalSum int64 = 0
	smallest := int64(arr[0])
	largest := int64(arr[0])

	// iterate over the number array and sum up all the values in the array
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

	// subtract the largest and smallest from totalSum to get back min and max sum
	minSum := totalSum - largest
	maxSum := totalSum - smallest
	return minSum, maxSum  // finally return the min and max values
}

func miniMaxSum(arr []int32) {
	minSum, maxSum := calculateMiniMaxSum(arr)
	fmt.Println(minSum, maxSum)
}
