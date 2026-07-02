package hackerrank

// func gradingStudents(grades []int32) []int32 {
// 	result := []int32{}

// 	for _, grade := range grades {
// 		nextMultiple := grade + (5 - grade%5)

// 		if grade >= 38 && nextMultiple-grade < 3 {
// 			result = append(result, nextMultiple)
// 		} else {
// 			result = append(result, grade)
// 		}
// 	}

// 	return result
// }

func gradingStudents(grades []int32) []int32 {
	result := []int32{}

	for _, grade := range grades {
		nextMultiple := grade + (5 - grade%5)

		if grade >= 38 && nextMultiple - grade < 3 {
			result = append(result, nextMultiple)
		} else {
			result = append(result, grade)
		}
	}

	return result 
}
