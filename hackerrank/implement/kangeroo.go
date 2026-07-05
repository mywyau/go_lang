package implement

/*
	Kangaroo problem notes

	We have two kangaroos moving along a number line.

	Kangaroo 1:
	- starts at position x1
	- jumps v1 meters each time

	Kangaroo 2:
	- starts at position x2
	- jumps v2 meters each time

	After n jumps:

		kangaroo 1 position = x1 + n*v1
		kangaroo 2 position = x2 + n*v2

	They meet if both positions are equal after the same number of jumps:

		x1 + n*v1 == x2 + n*v2

	For the usual HackerRank version, kangaroo 1 starts behind kangaroo 2:

		x1 < x2

	So kangaroo 1 must jump farther each time:

		v1 > v2

	If kangaroo 1 is not faster, it can never catch up.

	Once we know kangaroo 1 is faster, we check whether the distance gap
	can be closed exactly by the speed gap.

	Example:

		x1 = 0, v1 = 3
		x2 = 4, v2 = 2

		distanceGap = x2 - x1 = 4
		speedGap    = v1 - v2 = 1

		4 % 1 == 0, so they meet.

	Another example:

		x1 = 0, v1 = 4
		x2 = 5, v2 = 2

		distanceGap = 5
		speedGap    = 2

		5 % 2 != 0, so kangaroo 1 jumps past kangaroo 2
		but never lands on the same position at the same time.

	Time complexity:  O(1)
	Space complexity: O(1)
*/
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// If kangaroo 1 is slower or the same speed, it cannot catch kangaroo 2.
	//
	// This assumes the usual HackerRank setup where kangaroo 1 starts behind:
	//
	//	x1 < x2
	if v1 <= v2 {
		return "NO"
	}

	// The starting distance between the two kangaroos.
	distanceGap := x2 - x1

	// How much kangaroo 1 closes the gap by on each jump.
	speedGap := v1 - v2

	// If the distance gap divides evenly by the speed gap,
	// then kangaroo 1 lands exactly on kangaroo 2 after some number of jumps.
	//
	// If there is a remainder, kangaroo 1 may pass kangaroo 2,
	// but they will not land on the same position at the same time.
	if distanceGap%speedGap == 0 {
		return "YES"
	}

	return "NO"
}

// function with no notes

// func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
// 	if v1 <= v2 {
// 		return "NO"
// 	}

// 	distanceGap := x2 - x1
// 	speedGap := v1 - v2

// 	if distanceGap%speedGap == 0 {
// 		return "YES"
// 	}

// 	return "NO"
// }