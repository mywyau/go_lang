package searching

/*
LeetCode 69: Sqrt(x)

Given a non-negative integer x, return the square root of x rounded down
to the nearest integer.

The returned integer should be the largest integer r such that:

	r * r <= x

You must not use any built-in exponent function or square root function.

Examples:

Input:  x = 4
Output: 2

Input:  x = 8
Output: 2

Explanation:
The square root of 8 is 2.828..., and since we round down,
the answer is 2.

Approach:
Use binary search between 1 and x.

If mid * mid == x:
	return mid

If mid * mid < x:
	mid might be the answer, but there may be a larger valid answer,
	so search the right side.

If mid * mid > x:
	mid is too large, so search the left side.

When the loop finishes, right points to the largest valid integer
whose square is less than or equal to x.
*/

func MySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x

	for left <= right {
		mid := left + (right-left)/2
		square := mid * mid

		if square == x {
			return mid
		}

		if square < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return right
}