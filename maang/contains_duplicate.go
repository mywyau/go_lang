package maang


// naive approach
// func containsDuplicate(nums []int) bool {
// 	for i := 0; i < len(nums); i++ {
// 		for j := i + 1; j < len(nums); j++ {
// 			if nums[i] == nums[j] {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// more efficient hashmap

// func containsDuplicate(nums []int) bool {
// 	seen := make(map[int]bool)

// 	for _, num := range nums {
// 		if seen[num] {
// 			return true
// 		}

// 		seen[num] = true
// 	}

// 	return false
// }


// idiomatic
// func containsDuplicate(nums []int) bool {
// 	seen := make(map[int]struct{})

// 	for _, num := range nums {
// 		if _, exists := seen[num]; exists {
// 			return true
// 		}

// 		seen[num] = struct{}{}
// 	}

// 	return false
// }

func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = struct{}{}
	}

	return false
}
