// HackerRank: Dynamic Array
//
// What this problem is trying to teach:
//
// 1. 2D slices
//    We create a slice of slices:
//
//        seqList := make([][]int32, n)
//
//    This gives us n empty sequences.
//    Example:
//
//        n = 2
//        seqList = [[], []]
//
// 2. Dynamic appending
//    Each inner slice can grow using append:
//
//        seqList[idx] = append(seqList[idx], y)
//
//    This means:
//    "Go to sequence idx and add y to the end."
//
// 3. Query processing
//    Each query has 3 values:
//
//        queryType x y
//
//    There are two query types:
//
//        Type 1: append y to a selected sequence
//        Type 2: read a value from a selected sequence and update lastAnswer
//
// 4. lastAnswer
//    lastAnswer starts at 0.
//
//        var lastAnswer int32 = 0
//
//    It only changes when we process a type 2 query.
//
//    Type 2 queries also store lastAnswer in the answers slice,
//    because HackerRank wants us to return all values produced by type 2 queries.
//
// 5. XOR operator
//    In Go, ^ means bitwise XOR:
//
//        x ^ lastAnswer
//
//    For this problem, you mainly need to remember that HackerRank gives us this formula:
//
//        idx := (x ^ lastAnswer) % n
//
//    This formula decides which sequence we should use.
//
// 6. Modulo operator
//    The % operator keeps the index inside the valid range.
//
//    Example:
//
//        n = 2
//
//    Valid indexes are only:
//
//        0 and 1
//
//    So something % 2 will always give either 0 or 1.
//
// 7. Main idea
//    Type 1 queries write data.
//    Type 2 queries read data and update lastAnswer.
//
//    So the problem is mainly about managing state while processing queries.

func dynamicArray(n int32, queries [][]int32) []int32 {
	// Create n empty sequences.
	// Example: n = 2 gives [[], []]
	seqList := make([][]int32, n)

	// This value starts at 0 and changes after type 2 queries.
	var lastAnswer int32 = 0

	// We store the output from type 2 queries here.
	var answers []int32

	for _, query := range queries {
		queryType := query[0]
		x := query[1]
		y := query[2]

		// Choose which sequence to use.
		// The formula is given by HackerRank.
		idx := (x ^ lastAnswer) % n

		if queryType == 1 {
			// Add y to the selected sequence.
			seqList[idx] = append(seqList[idx], y)
		}

		if queryType == 2 {
			// Pick an item from the selected sequence.
			// We use modulo so the valueIndex stays inside the slice length.
			valueIndex := y % int32(len(seqList[idx]))

			// Update lastAnswer using the selected value.
			lastAnswer = seqList[idx][valueIndex]

			// Save lastAnswer because type 2 queries produce output.
			answers = append(answers, lastAnswer)
		}
	}

	return answers
}