package algorithms

/*
	Diagonal Difference

	Given a square matrix, calculate the absolute difference between
	the sum of the primary diagonal and the secondary diagonal.

	Primary diagonal:
	- Goes from top-left to bottom-right.
	- Index pattern: arr[i][i]

	Secondary diagonal:
	- Goes from top-right to bottom-left.
	- Index pattern: arr[i][n-1-i]

	Example 3x3 matrix indexes:

	arr[0][0]  arr[0][1]  arr[0][2]
	arr[1][0]  arr[1][1]  arr[1][2]
	arr[2][0]  arr[2][1]  arr[2][2]

	Primary diagonal:   arr[0][0], arr[1][1], arr[2][2]
	Secondary diagonal: arr[0][2], arr[1][1], arr[2][0]

	Finally, return the absolute difference between the two sums.
*/

// func diagonalDifference(arr [][]int32) int32 {
// 	var primaryDiagonal int32 = 0
// 	var secondaryDiagonal int32 = 0

// 	n := len(arr)

// 	for i := 0; i < n; i++ {
// 		primaryDiagonal += arr[i][i]
// 		secondaryDiagonal += arr[i][n-1-i]
// 	}

// 	difference := primaryDiagonal - secondaryDiagonal

// 	if difference < 0 {
// 		return -difference
// 	}

// 	return difference
// }

func diagonalDifference(arr [][]int32) int32 {
	var primaryDiagonal int32 = 0
	var secondaryDiagonal int32 = 0

	n := len(arr)

	for i := 0; i < n; i++ {
		primaryDiagonal += arr[i][i]
		secondaryDiagonal += arr[i][n-1-i]
	}

	difference := primaryDiagonal - secondaryDiagonal

	if difference < 0 {
		return -difference
	}

	return difference
}
