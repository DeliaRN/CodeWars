package kata15countvowels

import "strings"

func GetCount(str string) (count int) {
	count = 0
	vowels := "aeiouAEIOU"

	for _, char := range str {
		if strings.ContainsRune(vowels, char) {
			count++
		}

	}
	return count
}

func AlternativeWayOfCounting(str string) (count int) {
	for _, c := range str {
		switch c {
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return count
}
