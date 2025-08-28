package kata01unrepchar

func FirstNonRepeating(str string) string {

	charCount := make(map[rune]int)

	for _, char := range str {
		charCount[char]++
	}

	for _, char := range str {
		if charCount[char] == 1 {
			return string(char)
		}

}