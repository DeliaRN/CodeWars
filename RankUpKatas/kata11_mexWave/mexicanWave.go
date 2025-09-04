package kata11mexwave

import "strings"

func Wave(words string) []string {
	// Your code here and happy coding!
	var result []string

	for i, letter := range words {
		if letter == ' ' {
			continue
		}

		// Build a new string: lowercase prefix + uppercase char + rest
		//i=2 --> "he"(i0 to i2 -> h, e) + "L" (letter at i(2) uppercased) + "lo" (i(3) to end)
		wave := words[:i] + strings.ToUpper(string(letter)) + words[i+1:]
		result = append(result, wave)

	}

	return result

}
