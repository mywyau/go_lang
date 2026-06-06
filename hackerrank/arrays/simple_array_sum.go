package arrays

/*
   Problem: Simple Array Sum

   We are given a list of numbers.
   We need to add all the numbers together.
   Then we return the final total.

   Problem solving steps:
   1. Create a variable called sum and start it at 0.
   2. Loop through every number in the array.
   3. Add the current number to sum.
   4. After the loop finishes, return sum.

   Example:
   ar = [1, 2, 3]

   Start:
   sum = 0

   Add 1:
   sum = 1

   Add 2:
   sum = 3

   Add 3:
   sum = 6

   Return 6.
*/

func simpleArraySum(nums []int32) int32 {
	var sum int32 = 0
	for _, v := range nums {
		sum += v
	}
	return sum
}

// func simpleArraySum(ar []int32) int32 {
// 	var sum int32 = 0

// 	for _, v := range ar {
// 		sum += v
// 	}

// 	return sum
// }
