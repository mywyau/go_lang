package arrays

// import "fmt"

// we are trying to learn to loop through 2d arrays and apply a pattern to the elements

// func hourglassSum(arr [][]int32) int32 {

// 	var maxSum int32 = -999999

// 	for row := 0; row <= 3; row++ {
// 		for col := 0; col <= 3; col++ {

// 			top := arr[row][col] + arr[row][col+1] + arr[row][col+2]

// 			middle := arr[row+1][col+1]

// 			bottom := arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]

// 			sum := top + middle + bottom

// 			if sum > maxSum {
// 				maxSum = sum
// 			}
// 		}
// 	}

// 	return maxSum
// }

// func hourglassSum(arr [][]int32) int32 {

// 	var maxSum int32 = -999999

// 	for row := 0; row <= 3; row++ {
// 		for col := 0; col <= 3; col++ {
// 			top := arr[row][col] + arr[row][col+1] + arr[row][col+2]
// 			middle := arr[row+1][col+1]
// 			bottom := arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]
// 			sum := top + middle + bottom

// 			if sum > maxSum {
// 				maxSum = sum
// 			}
// 		}
// 	}
// 	return maxSum
// }

func hourglassSum(arr [][]int32) int32 {

	var maxSum int32 = -999999

	for row := 0; row <= 3; row++ {
		for col := 0; col <= 3; col++ {

			top := arr[row][col] + arr[row][col+1] + arr[row][col+2]
			middle := arr[row+1][col+1]
			bottom := arr[row+2][col] + arr[row+2][col+1] + arr[row+2][col+2]
			sum := top + middle + bottom

			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

// func main() {
// 	arr := [][]int{
// 		{1, 1, 1, 0, 0, 0},
// 		{0, 1, 0, 0, 0, 0},
// 		{1, 1, 1, 0, 0, 0},
// 		{0, 0, 2, 4, 4, 0},
// 		{0, 0, 0, 2, 0, 0},
// 		{0, 0, 1, 2, 4, 0},
// 	}

// 	fmt.Println(hourglassSum(arr)) // 19
// }
