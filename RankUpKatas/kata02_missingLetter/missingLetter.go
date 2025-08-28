package kata02missingletter

import "strings"

func FindMissingLetter(chars []rune) rune {
	var alphabet string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var alphabetRunes []rune = []rune(alphabet)
	for i, char := range chars {
		alphPosition := strings.IndexRune(alphabet, char)
		if chars[i+1] != alphabetRunes[alphPosition+1] {
			return alphabetRunes[alphPosition+1]
		}
	}
	return 0
}

func BETTERApproach(chars []rune) rune {
	expected := chars[0]
	for _, actual := range chars[1:] {
		expected++ // move to the next letter
		if actual != expected {
			break // found the gap
		}
	}
	return expected // this is the missing letter
}

//EXPLANATION:
// a rune is an alias for int32, which means it represents
// a Unicode code point as a number.
//When you do expected++, you’re actually incrementing
// the numeric value of the rune.
// For example, if expected is 'a', its numeric value is 97.
// After expected++, it becomes 98, which is 'b'.
