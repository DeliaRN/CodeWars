package kata01unrepchar

import "strings"

func FirstNonRepeating(str string) string {

	charCount := make(map[rune]int)
	lowerStr := strings.ToLower(str)

	for _, char := range lowerStr {
		charCount[char]++
	}

	for i, char := range str {
		if charCount[rune(lowerStr[i])] == 1 {
			return string(char)
		}
	}
	return ""
}
