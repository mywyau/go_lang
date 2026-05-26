from pathlib import Path

import "fmt"

/*
Go Maps Learning File
=====================

A map in Go is like a dictionary / hash map.

It stores key-value pairs.

Examples:
- string -> int
- int -> bool
- string -> []string

Maps are very important for coding interviews because they let you
look things up quickly.

Common uses:
1. Counting frequencies
2. Checking if something exists
3. Storing indexes
4. Building sets
5. Grouping values
*/

// ------------------------------------------------------------
// 1. Creating maps
// ------------------------------------------------------------

func creatingMaps() {
	// Option 1: map literal
	ages := map[string]int{
		"Michael": 30,
		"Sarah":   25,
		"John":    40,
	}

	fmt.Println("Ages:", ages)

	// Option 2: make
	scores := make(map[string]int)

	scores["Alice"] = 90
	scores["Bob"] = 75

	fmt.Println("Scores:", scores)
}

// ------------------------------------------------------------
// 2. Reading from maps
// ------------------------------------------------------------

func readingMaps() {
	ages := map[string]int{
		"Michael": 30,
		"Sarah":   25,
	}

	fmt.Println("Michael age:", ages["Michael"])

	// If a key does not exist, Go returns the zero value.
	// For int, the zero value is 0.
	fmt.Println("Unknown age:", ages["Unknown"])
}

// ------------------------------------------------------------
// 3. Checking if a key exists
// ------------------------------------------------------------

func checkingExists() {
	ages := map[string]int{
		"Michael": 30,
	}

	age, exists := ages["Michael"]

	if exists {
		fmt.Println("Found Michael:", age)
	} else {
		fmt.Println("Michael not found")
	}

	unknownAge, unknownExists := ages["Unknown"]

	if unknownExists {
		fmt.Println("Found Unknown:", unknownAge)
	} else {
		fmt.Println("Unknown not found")
	}
}

// ------------------------------------------------------------
// 4. Updating and deleting values
// ------------------------------------------------------------

func updatingAndDeleting() {
	ages := map[string]int{
		"Michael": 30,
	}

	// Add new key
	ages["Sarah"] = 25

	// Update existing key
	ages["Michael"] = 31

	fmt.Println("After add/update:", ages)

	// Delete key
	delete(ages, "Sarah")

	fmt.Println("After delete:", ages)
}

// ------------------------------------------------------------
// 5. Looping over a map
// ------------------------------------------------------------

func loopingMap() {
	ages := map[string]int{
		"Michael": 30,
		"Sarah":   25,
		"John":    40,
	}

	for name, age := range ages {
		fmt.Println(name, age)
	}

	// Note:
	// Map iteration order is not guaranteed.
	// So the printed order may be different each time.
}

// ------------------------------------------------------------
// 6. Counting frequency
// ------------------------------------------------------------

func countNumbers(nums []int) map[int]int {
	counts := make(map[int]int)

	for _, num := range nums {
		counts[num]++
	}

	return counts
}

// ------------------------------------------------------------
// 7. Using a map as a set
// ------------------------------------------------------------

func containsDuplicate(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}

		seen[num] = true
	}

	return false
}

// A slightly more memory-efficient version uses map[int]struct{}.
// struct{} takes no extra memory for the value.

func containsDuplicateWithStructSet(nums []int) bool {
	seen := make(map[int]struct{})

	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}

		seen[num] = struct{}{}
	}

	return false
}

// ------------------------------------------------------------
// 8. Two Sum
// ------------------------------------------------------------

/*
Problem:

Given nums and a target, return the indexes of two numbers
that add up to target.

Example:

nums = []int{2, 7, 11, 15}
target = 9

2 + 7 = 9

Return []int{0, 1}
*/

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, num := range nums {
		needed := target - num

		if index, exists := seen[needed]; exists {
			return []int{index, i}
		}

		seen[num] = i
	}

	return []int{}
}

// ------------------------------------------------------------
// 9. First repeated number
// ------------------------------------------------------------

func firstRepeated(nums []int) int {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return num
		}

		seen[num] = true
	}

	return -1
}

// ------------------------------------------------------------
// 10. Character frequency
// ------------------------------------------------------------

func charFrequency(s string) map[rune]int {
	counts := make(map[rune]int)

	for _, ch := range s {
		counts[ch]++
	}

	return counts
}

// ------------------------------------------------------------
// 11. Group words by length
// ------------------------------------------------------------

func groupWordsByLength(words []string) map[int][]string {
	groups := make(map[int][]string)

	for _, word := range words {
		length := len(word)
		groups[length] = append(groups[length], word)
	}

	return groups
}

// ------------------------------------------------------------
// Main function to run examples
// ------------------------------------------------------------

func main() {
	fmt.Println("=== Creating maps ===")
	creatingMaps()

	fmt.Println("\n=== Reading maps ===")
	readingMaps()

	fmt.Println("\n=== Checking exists ===")
	checkingExists()

	fmt.Println("\n=== Updating and deleting ===")
	updatingAndDeleting()

	fmt.Println("\n=== Looping map ===")
	loopingMap()

	fmt.Println("\n=== Count numbers ===")
	fmt.Println(countNumbers([]int{1, 2, 2, 3, 3, 3}))

	fmt.Println("\n=== Contains duplicate ===")
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))

	fmt.Println("\n=== Contains duplicate with struct set ===")
	fmt.Println(containsDuplicateWithStructSet([]int{1, 2, 3, 1}))
	fmt.Println(containsDuplicateWithStructSet([]int{1, 2, 3, 4}))

	fmt.Println("\n=== Two Sum ===")
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))

	fmt.Println("\n=== First repeated ===")
	fmt.Println(firstRepeated([]int{5, 1, 2, 5, 3}))
	fmt.Println(firstRepeated([]int{1, 2, 3, 4}))

	fmt.Println("\n=== Character frequency ===")
	fmt.Println(charFrequency("hello"))

	fmt.Println("\n=== Group words by length ===")
	fmt.Println(groupWordsByLength([]string{"go", "map", "code", "tree", "hi"}))
}

/*
Practice Tasks
==============

Try writing these yourself without looking at the answers above.

1. Count numbers
----------------

Write:

func countNumbers(nums []int) map[int]int

Example:

countNumbers([]int{1, 2, 2, 3, 3, 3})

Should return:

map[int]int{
	1: 1,
	2: 2,
	3: 3,
}

2. Contains duplicate
---------------------

Write:

func containsDuplicate(nums []int) bool

Examples:

containsDuplicate([]int{1, 2, 3, 1}) // true
containsDuplicate([]int{1, 2, 3, 4}) // false

3. Two Sum
----------

Write:

func twoSum(nums []int, target int) []int

Example:

twoSum([]int{2, 7, 11, 15}, 9)

Should return:

[]int{0, 1}

4. First repeated number
------------------------

Write:

func firstRepeated(nums []int) int

Example:

firstRepeated([]int{5, 1, 2, 5, 3})

Should return:

5

If no number repeats, return -1.

5. Character frequency
----------------------

Write:

func charFrequency(s string) map[rune]int

Example:

charFrequency("hello")

Should return counts like:

h: 1
e: 1
l: 2
o: 1

Key interview patterns
======================

Map as frequency counter:

counts := make(map[int]int)

for _, num := range nums {
	counts[num]++
}

Map as set:

seen := make(map[int]bool)

for _, num := range nums {
	if seen[num] {
		return true
	}

	seen[num] = true
}

Map for indexes:

seen := make(map[int]int)

for i, num := range nums {
	seen[num] = i
}

Check key exists:

value, exists := myMap[key]

if exists {
	fmt.Println(value)
}
*/
'''

path = Path("/mnt/data/maps_learning.go")
path.write_text(content)

print(f"Created file: {path}")