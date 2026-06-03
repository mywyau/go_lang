package arrays

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


