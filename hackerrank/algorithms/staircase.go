package algorithms

import (
	"fmt"
	"strings"
)

/*
	Staircase

	Given a number n, print a right-aligned staircase of height n.

	Example when n = 6:

	     #
	    ##
	   ###
	  ####
	 #####
	######

	Each row contains:
	- some spaces first
	- then some # symbols

	For each row i from 1 to n:

	    spaces = n - i
	    hashes = i

	So when n = 6:

	    row 1 -> 5 spaces + 1 #
	    row 2 -> 4 spaces + 2 #
	    row 3 -> 3 spaces + 3 #
	    row 4 -> 2 spaces + 4 #
	    row 5 -> 1 space  + 5 #
	    row 6 -> 0 spaces + 6 #
*/

func staircase(n int32) {
	for i := int32(1); i <= n; i++ {
		spaces := n - i
		hashes := i

		fmt.Print(strings.Repeat(" ", int(spaces)))
		fmt.Print(strings.Repeat("#", int(hashes)))
		fmt.Println()
	}
}

// testable helper function

func staircaseOutput(n int32) string {
	var result strings.Builder

	for i := int32(1); i <= n; i++ {
		spaces := n - i
		hashes := i

		result.WriteString(strings.Repeat(" ", int(spaces)))
		result.WriteString(strings.Repeat("#", int(hashes)))
		result.WriteString("\n")
	}

	return result.String()
}
