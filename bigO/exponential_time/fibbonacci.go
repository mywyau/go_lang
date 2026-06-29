package exponentialtime

// Exponential time
// O(2^n)
// Common with naive recursion where each call branches into two more calls.
// This becomes slow very quickly.


func fib(n int) int {

	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
