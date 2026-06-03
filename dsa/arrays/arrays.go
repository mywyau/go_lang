package arrays

import "fmt"

// ARRAYS
//
// An array has a fixed size.
// The size is part of the type.
//
// [3]int means:
// - array
// - length 3
// - stores int values
var arr [3]int = [3]int{1, 2, 3}

func ArrayExamples() {
	// Create an array with 3 numbers
	numbers := [3]int{10, 20, 30}

	fmt.Println("array:", numbers)
	fmt.Println("first item:", numbers[0])
	fmt.Println("second item:", numbers[1])
	fmt.Println("third item:", numbers[2])

	// Change an item
	numbers[1] = 99
	fmt.Println("after update:", numbers)

	// Length of an array
	fmt.Println("array length:", len(numbers))

	// Loop through an array
	for i := 0; i < len(numbers); i++ {
		fmt.Println("index:", i, "value:", numbers[i])
	}

	// Cleaner loop using range
	for index, value := range numbers {
		fmt.Println("range index:", index, "range value:", value)
	}
}

// SLICES
//
// A slice is like a flexible array.
// It can grow and shrink.
//
// []int means:
// - slice
// - stores int values
func SliceExamples() {
	nums := []int{1, 2, 3}

	fmt.Println("slice:", nums)

	// Add to a slice using append
	nums = append(nums, 4)
	nums = append(nums, 5)

	fmt.Println("after append:", nums)

	// Access by index
	fmt.Println("first item:", nums[0])
	fmt.Println("last item:", nums[len(nums)-1])

	// Update an item
	nums[0] = 100
	fmt.Println("after update:", nums)

	// Length of a slice
	fmt.Println("slice length:", len(nums))

	// Loop through a slice
	for i := 0; i < len(nums); i++ {
		fmt.Println("index:", i, "value:", nums[i])
	}

	// Cleaner loop using range
	for index, value := range nums {
		fmt.Println("range index:", index, "range value:", value)
	}
}

func SliceCuttingExamples() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("original:", nums)

	// Take from index 1 up to, but not including, index 4
	fmt.Println("nums[1:4]:", nums[1:4]) // [20 30 40]

	// Take from the start up to index 3
	fmt.Println("nums[:3]:", nums[:3]) // [10 20 30]

	// Take from index 2 to the end
	fmt.Println("nums[2:]:", nums[2:]) // [30 40 50]

	// Take the whole slice
	fmt.Println("nums[:]:", nums[:]) // [10 20 30 40 50]
}

func SliceMakeExamples() {
	// make creates a slice with a starting length
	nums := make([]int, 3)

	fmt.Println("made slice:", nums) // [0 0 0]

	nums[0] = 10
	nums[1] = 20
	nums[2] = 30

	fmt.Println("after updates:", nums)

	// make with length and capacity
	scores := make([]int, 0, 5)

	fmt.Println("scores:", scores)
	fmt.Println("length:", len(scores))
	fmt.Println("capacity:", cap(scores))

	scores = append(scores, 100)
	scores = append(scores, 200)

	fmt.Println("after append:", scores)
	fmt.Println("length:", len(scores))
	fmt.Println("capacity:", cap(scores))
}

func ArrayVsSliceExamples() {
	// Array: fixed size
	arrayNumbers := [3]int{1, 2, 3}

	// Slice: flexible size
	sliceNumbers := []int{1, 2, 3}

	fmt.Println("array:", arrayNumbers)
	fmt.Println("slice:", sliceNumbers)

	// This would NOT work because arrays cannot grow:
	// arrayNumbers = append(arrayNumbers, 4)

	// This works because slices can grow:
	sliceNumbers = append(sliceNumbers, 4)

	fmt.Println("slice after append:", sliceNumbers)
}