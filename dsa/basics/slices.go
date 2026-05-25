package basics

import "fmt"

/*
Go Slices Learning File
=======================

A slice is Go's flexible, commonly-used list type.

You will use slices constantly in:
- coding interviews
- LeetCode problems
- real Go applications

Mental model:

Array:
- fixed size
- size is part of the type
- example: [3]int{1, 2, 3}

Slice:
- flexible size
- built on top of an array
- example: []int{1, 2, 3}

Most of the time, use slices instead of arrays.
*/

// ------------------------------------------------------------
// 1. Arrays vs slices
// ------------------------------------------------------------

func arraysVsSlices() {
	// Array: fixed size
	arr := [3]int{1, 2, 3}

	// Slice: flexible size
	nums := []int{1, 2, 3}

	fmt.Println("Array:", arr)
	fmt.Println("Slice:", nums)

	// You can append to a slice
	nums = append(nums, 4)

	fmt.Println("After append:", nums)

	// You cannot append directly to an array variable.
	// Arrays have fixed size.
}

// ------------------------------------------------------------
// 2. Creating slices
// ------------------------------------------------------------

func creatingSlices() {
	// Option 1: slice literal
	nums := []int{10, 20, 30}

	fmt.Println("nums:", nums)

	// Option 2: make
	// This creates a slice with length 3.
	values := make([]int, 3)

	values[0] = 100
	values[1] = 200
	values[2] = 300

	fmt.Println("values:", values)

	// Option 3: make with length and capacity
	// length = number of usable elements
	// capacity = how much space exists before Go may need a bigger backing array
	withCapacity := make([]int, 0, 5)

	fmt.Println("withCapacity:", withCapacity)
	fmt.Println("len:", len(withCapacity))
	fmt.Println("cap:", cap(withCapacity))
}

// ------------------------------------------------------------
// 3. len and cap
// ------------------------------------------------------------

func lenAndCap() {
	nums := make([]int, 0, 3)

	fmt.Println("start:", nums, "len:", len(nums), "cap:", cap(nums))

	nums = append(nums, 1)
	fmt.Println("after 1:", nums, "len:", len(nums), "cap:", cap(nums))

	nums = append(nums, 2)
	fmt.Println("after 2:", nums, "len:", len(nums), "cap:", cap(nums))

	nums = append(nums, 3)
	fmt.Println("after 3:", nums, "len:", len(nums), "cap:", cap(nums))

	// This may cause Go to allocate a larger backing array.
	nums = append(nums, 4)
	fmt.Println("after 4:", nums, "len:", len(nums), "cap:", cap(nums))
}

// ------------------------------------------------------------
// 4. Accessing and updating values
// ------------------------------------------------------------

func accessingAndUpdating() {
	nums := []int{10, 20, 30}

	fmt.Println("first:", nums[0])
	fmt.Println("second:", nums[1])

	nums[1] = 99

	fmt.Println("after update:", nums)
}

// ------------------------------------------------------------
// 5. Looping over slices
// ------------------------------------------------------------

func loopingSlices() {
	nums := []int{10, 20, 30}

	// Normal for loop
	for i := 0; i < len(nums); i++ {
		fmt.Printf("normal loop: index=%d value=%d\n", i, nums[i])
	}

	// Range loop
	for i, value := range nums {
		fmt.Printf("range loop: index=%d value=%d\n", i, value)
	}

	// If you do not need the index, use _
	for _, value := range nums {
		fmt.Printf("value only: %d\n", value)
	}
}

// ------------------------------------------------------------
// 6. Slicing a slice
// ------------------------------------------------------------

func slicingSlices() {
	nums := []int{10, 20, 30, 40, 50}

	fmt.Println("original:", nums)

	// nums[start:end]
	// start is included
	// end is excluded
	fmt.Println("nums[1:4]:", nums[1:4]) // 20, 30, 40

	fmt.Println("nums[:3]:", nums[:3])   // 10, 20, 30
	fmt.Println("nums[2:]:", nums[2:])   // 30, 40, 50
	fmt.Println("nums[:]:", nums[:])     // full slice
}

// ------------------------------------------------------------
// 7. Important slice gotcha: shared backing array
// ------------------------------------------------------------

func sharedBackingArray() {
	nums := []int{10, 20, 30, 40, 50}

	part := nums[1:4]

	fmt.Println("nums before:", nums)
	fmt.Println("part before:", part)

	// Updating part also updates nums because they share the same backing array.
	part[0] = 999

	fmt.Println("nums after:", nums)
	fmt.Println("part after:", part)
}

// ------------------------------------------------------------
// 8. Copying slices
// ------------------------------------------------------------

func copyingSlices() {
	original := []int{1, 2, 3}

	copySlice := make([]int, len(original))

	copy(copySlice, original)

	copySlice[0] = 999

	fmt.Println("original:", original)
	fmt.Println("copySlice:", copySlice)
}

// ------------------------------------------------------------
// 9. Nil slice vs empty slice
// ------------------------------------------------------------

func nilVsEmptySlice() {
	var nilSlice []int

	emptySlice := []int{}

	fmt.Println("nilSlice:", nilSlice, "len:", len(nilSlice), "is nil:", nilSlice == nil)
	fmt.Println("emptySlice:", emptySlice, "len:", len(emptySlice), "is nil:", emptySlice == nil)

	// Both can be appended to.
	nilSlice = append(nilSlice, 1)
	emptySlice = append(emptySlice, 1)

	fmt.Println("nilSlice after append:", nilSlice)
	fmt.Println("emptySlice after append:", emptySlice)
}

// ------------------------------------------------------------
// 10. Passing slices to functions
// ------------------------------------------------------------

func mutateFirst(nums []int) {
	if len(nums) == 0 {
		return
	}

	// This mutates the underlying array, so the caller sees the change.
	nums[0] = 999
}

func passingSlicesToFunctions() {
	nums := []int{1, 2, 3}

	fmt.Println("before:", nums)

	mutateFirst(nums)

	fmt.Println("after:", nums)
}

// ------------------------------------------------------------
// 11. Sum slice
// ------------------------------------------------------------

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

// ------------------------------------------------------------
// 12. Find max
// ------------------------------------------------------------

func max(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	best := nums[0]

	for _, num := range nums {
		if num > best {
			best = num
		}
	}

	return best
}

// ------------------------------------------------------------
// 13. Reverse slice
// ------------------------------------------------------------

func reverse(nums []int) []int {
	left := 0
	right := len(nums) - 1

	for left < right {
		nums[left], nums[right] = nums[right], nums[left]

		left++
		right--
	}

	return nums
}

// ------------------------------------------------------------
// 14. Filter even numbers
// ------------------------------------------------------------

func filterEven(nums []int) []int {
	result := []int{}

	for _, num := range nums {
		if num%2 == 0 {
			result = append(result, num)
		}
	}

	return result
}

// ------------------------------------------------------------
// 15. Contains value
// ------------------------------------------------------------

func contains(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}

	return false
}

// ------------------------------------------------------------
// 16. Remove value
// ------------------------------------------------------------

func removeValue(nums []int, target int) []int {
	result := []int{}

	for _, num := range nums {
		if num != target {
			result = append(result, num)
		}
	}

	return result
}

// ------------------------------------------------------------
// 17. Two pointers: reverse string-like rune slice
// ------------------------------------------------------------

func reverseRunes(chars []rune) []rune {
	left := 0
	right := len(chars) - 1

	for left < right {
		chars[left], chars[right] = chars[right], chars[left]

		left++
		right--
	}

	return chars
}

// ------------------------------------------------------------
// 18. Move zeroes to the end
// ------------------------------------------------------------

/*
Problem:

Given a slice of numbers, move all zeroes to the end while keeping
the order of the non-zero numbers.

Example:

Input:  []int{0, 1, 0, 3, 12}
Output: []int{1, 3, 12, 0, 0}
*/

func moveZeroes(nums []int) []int {
	insertPosition := 0

	// First, move all non-zero numbers forward.
	for _, num := range nums {
		if num != 0 {
			nums[insertPosition] = num
			insertPosition++
		}
	}

	// Then fill the rest with zeroes.
	for insertPosition < len(nums) {
		nums[insertPosition] = 0
		insertPosition++
	}

	return nums
}

// ------------------------------------------------------------
// Main function to run examples
// ------------------------------------------------------------

func main() {
	fmt.Println("=== Arrays vs slices ===")
	arraysVsSlices()

	fmt.Println("\n=== Creating slices ===")
	creatingSlices()

	fmt.Println("\n=== len and cap ===")
	lenAndCap()

	fmt.Println("\n=== Accessing and updating ===")
	accessingAndUpdating()

	fmt.Println("\n=== Looping slices ===")
	loopingSlices()

	fmt.Println("\n=== Slicing slices ===")
	slicingSlices()

	fmt.Println("\n=== Shared backing array ===")
	sharedBackingArray()

	fmt.Println("\n=== Copying slices ===")
	copyingSlices()

	fmt.Println("\n=== Nil vs empty slice ===")
	nilVsEmptySlice()

	fmt.Println("\n=== Passing slices to functions ===")
	passingSlicesToFunctions()

	fmt.Println("\n=== Sum ===")
	fmt.Println(sum([]int{1, 2, 3, 4}))

	fmt.Println("\n=== Max ===")
	fmt.Println(max([]int{4, 10, 2, 8}))

	fmt.Println("\n=== Reverse ===")
	fmt.Println(reverse([]int{1, 2, 3, 4, 5}))

	fmt.Println("\n=== Filter even ===")
	fmt.Println(filterEven([]int{1, 2, 3, 4, 5, 6}))

	fmt.Println("\n=== Contains ===")
	fmt.Println(contains([]int{1, 2, 3}, 2))
	fmt.Println(contains([]int{1, 2, 3}, 99))

	fmt.Println("\n=== Remove value ===")
	fmt.Println(removeValue([]int{1, 2, 3, 2, 4}, 2))

	fmt.Println("\n=== Reverse runes ===")
	fmt.Println(string(reverseRunes([]rune("hello"))))
	fmt.Println(string(reverseRunes([]rune("你好"))))

	fmt.Println("\n=== Move zeroes ===")
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12}))
}

/*
Practice Tasks
==============

Try writing these yourself without looking at the answers above.

1. Sum a slice
--------------

Write:

func sum(nums []int) int

Example:

sum([]int{1, 2, 3, 4}) // 10

2. Find max
-----------

Write:

func max(nums []int) int

Example:

max([]int{4, 10, 2, 8}) // 10

Think about what you should return for an empty slice.

3. Reverse a slice
------------------

Write:

func reverse(nums []int) []int

Example:

reverse([]int{1, 2, 3, 4}) // []int{4, 3, 2, 1}

Hint:
Use two pointers:
- left starts at 0
- right starts at len(nums) - 1
- swap values while left < right

4. Filter even numbers
----------------------

Write:

func filterEven(nums []int) []int

Example:

filterEven([]int{1, 2, 3, 4, 5, 6}) // []int{2, 4, 6}

Hint:
Create a result slice and append values that match the condition.

5. Contains value
-----------------

Write:

func contains(nums []int, target int) bool

Example:

contains([]int{1, 2, 3}, 2)  // true
contains([]int{1, 2, 3}, 99) // false

6. Remove value
---------------

Write:

func removeValue(nums []int, target int) []int

Example:

removeValue([]int{1, 2, 3, 2, 4}, 2) // []int{1, 3, 4}

7. Move zeroes
--------------

Write:

func moveZeroes(nums []int) []int

Example:

moveZeroes([]int{0, 1, 0, 3, 12}) // []int{1, 3, 12, 0, 0}

Key slice patterns
==================

Create a slice:

nums := []int{1, 2, 3}

Append to a slice:

nums = append(nums, 4)

Loop over a slice:

for i, value := range nums {
	fmt.Println(i, value)
}

Slice part of a slice:

part := nums[1:3]

Copy a slice:

copySlice := make([]int, len(nums))
copy(copySlice, nums)

Two pointer pattern:

left := 0
right := len(nums) - 1

for left < right {
	nums[left], nums[right] = nums[right], nums[left]
	left++
	right--
}

Important notes
===============

1. len(nums) gives the number of elements.

2. cap(nums) gives the capacity of the slice before Go may need
   to allocate a bigger backing array.

3. append can return a new slice, so always assign it:

   nums = append(nums, 10)

4. Slices can share the same backing array.

   If you do:

   part := nums[1:3]

   then changing part may also change nums.

5. For most DSA problems, use slices rather than arrays.
*/
'''

path = Path("/mnt/data/slices_learning.go")
path.write_text(content)

print(f"Created file: {path}")