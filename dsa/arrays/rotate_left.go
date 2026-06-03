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

func rotateLeft(d int32, arr []int32) []int32 {
	n := len(arr)

	if n == 0 {
		return arr
	}

	// Convert d from int32 to int because slice indexes need int
	rotations := int(d) % n

	result := []int32{}

	result = append(result, arr[rotations:]...)
	result = append(result, arr[:rotations]...)

	return result
}

func rotateLeft2(d int32, arr []int32) []int32 {
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
