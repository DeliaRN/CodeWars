package kata03camelcase

import "strings"

func CamelCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	words := strings.Split(s, " ")
	for i, word := range words {
		if len(word) > 0 {
			words[i] = strings.ToUpper(string(word[0])) + word[1:]
		}

	}
	return strings.Join(words, "")

}
