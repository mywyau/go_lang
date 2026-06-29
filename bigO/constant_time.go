package bigo


// this is constant time 

// Absolutely. Big O is about answering:
// As the input gets bigger, how does the work grow?
// It does not usually mean exact speed in seconds. It means the growth pattern.

// this will always be O(1) since it just accesses a single element 
// even with 1, 10, 100, 1000, 10000, 100000, 1000000 etc. 

func firfirstElement(nums []int) int {
	return nums[0]
}





