package maang

/*
Anagram

Two words are anagrams if their letters can be rearranged
to form the other word.

Given a string, split it into two contiguous substrings of
equal length.

Determine the minimum number of characters that must be
changed to make the two substrings anagrams of one another.

Return:
  - the minimum number of character changes required
  - -1 if the string has an odd length and cannot be split
    into two equal halves

Example:

Input:
    "abccde"

Split:
    "abc" and "cde"

The character 'c' already matches.

Change:
    'a' -> 'd'
    'b' -> 'e'

Two changes are required, so the result is 2.

The input contains only lowercase English letters from
'a' to 'z'.
*/

func SubstringAnagram(s string) int32 {
	if len(s)%2 != 0 {
		return -1
	}

	middle := len(s) / 2

	var counts [26]int

	// Count the letters in the first half.
	for i := 0; i < middle; i++ {
		index := s[i] - 'a'
		counts[index]++
	}

	// Cancel matching letters from the second half.
	for i := middle; i < len(s); i++ {
		index := s[i] - 'a'
		counts[index]--
	}

	var changes int32

	// Positive values represent unmatched letters
	// remaining in the first half.
	for _, count := range counts {
		if count > 0 {
			changes += int32(count)
		}
	}

	return changes
}
