package arrays 

func dynamicArray(n int32, queries [][]int32) []int32 {
	sequences := make([][]int32, n)

	var lastAnswer int32
	var result []int32

	for _, query := range queries {
		queryType := query[0]
		x := query[1]
		y := query[2]

		idx := (x ^ lastAnswer) % n

		if queryType == 1 {
			sequences[idx] = append(sequences[idx], y)
		} else {
			seq := sequences[idx]
			lastAnswer = seq[y%int32(len(seq))]
			result = append(result, lastAnswer)
		}
	}

	return result
}


