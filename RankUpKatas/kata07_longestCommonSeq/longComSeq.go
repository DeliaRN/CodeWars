package kata07longestcommonseq

import "strings"

func LCS(s1, s2 string) string {

	if len(s1)*len(s2) == 0 {
		return ""
	}
	if s1[0] == s2[0] {
		return s1[:1] + LCS(s1[1:], s2[1:])
	}
	w1, w2 := LCS(s1, s2[1:]), LCS(s1[1:], s2)
	if len(w2) > len(w1) {
		return w2
	}
	return w1
}

// Explanation: If either string is empty, return an empty string.
// If the first characters of both strings match, include that character
// in the result and recursively call LCS on the remaining substrings.
// If they don't match, recursively call LCS twice:
// once excluding the first character of s2
// and once excluding the first character of s1.
// Return the longer of the two results.

func LCSnotWorkingWell(s1, s2 string) string {
	var result string = ""
	for _, char := range s1 {
		//For each character in s2, check if it exists in s1 and add it once to result if it does
		if strings.ContainsRune(s2, char) && !strings.ContainsRune(result, char) {
			result += string(char)
		}
	}

	return result
}
