package arrays

// func rotateLeft(d int32, arr []int32) []int32 {
// 	// Get the length of the slice
// 	n := len(arr)

// 	// If the slice is empty, there is nothing to rotate
// 	if n == 0 {
// 		return arr
// 	}

// 	// Convert d from int32 to int because Go slice indexes use int
// 	// Use % n so large rotations wrap around
// 	// Example: d = 7, n = 5, so 7 % 5 = 2
// 	rotations := int(d) % n

// 	// Create an empty result slice
// 	// This will store the rotated answer
// 	result := []int32{}

// 	// Take everything from the rotation point to the end
// 	// Example:
// 	// arr = [1 2 3 4 5], rotations = 2
// 	// arr[2:] gives [3 4 5]
// 	result = append(result, arr[rotations:]...)

// 	// Take everything before the rotation point
// 	// arr[:2] gives [1 2]
// 	// Append it after [3 4 5]
// 	result = append(result, arr[:rotations]...)

// 	// Return the rotated slice
// 	// Final example result: [3 4 5 1 2]
// 	return result
// }

/*
	Main idea:

	A left rotation moves the first d items to the end of the array.

	Example:
	arr = [1, 2, 3, 4, 5]
	d = 2

	Move the first 2 items to the end:

	[3, 4, 5] + [1, 2] = [3, 4, 5, 1, 2]

	We use d % len(arr) so extra full rotations do not matter.
	For example, rotating 5 items by 5 gives the same array.
*/

// func rotateLeft(d int32, arr []int32) []int32 {
// 	n := len(arr)

// 	if n == 0 {
// 		return arr
// 	}

// 	// Convert d from int32 to int because slice indexes need int
// 	// we use modulo because for larger number of rotations which are greater than the array ends up resetting the array,
// 	//  and we need to find the remainder == number of actual rotations
// 	rotations := int(d) % n

// 	result := []int32{} // empty slice to build the answer

// 	// The ... in Go is like saying: take all the items from this slice and append them one by one.
// 	//  e.g. arr = [1, 2, 3, 4, 5],  d = 2 this would append [3,4,5] to the empty result
// 	result = append(result, arr[rotations:]...)

// 	// then we append [1,2] to [3,4,5] = [3,4,5,1,2]
// 	result = append(result, arr[:rotations]...)

// 	return result // which results in the array rotated to the left by 2
// }

func rotateLeft(d int32, arr []int32) []int32 {
	n := len(arr)

	if n == 0 {
		return arr
	}

	rotations := int(d) % n

	result := []int32{}
	result = append(result, arr[rotations:]...)
	result = append(result, arr[:rotations]...)

	return result
}
