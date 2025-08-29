package kata05vigenerecipherhelper

import "strings"

type VigenèreCipher struct {
	Key   string
	Alpha string
}

func (c VigenèreCipher) Encode(str string) string {
	// c.Key = "PASSWORD"
	// c.Alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// str = "HELLO"

	KeyRunes := []rune(c.Key)            // KeyRunes = ['P','A','S','S','W','O','R','D']
	alphabetRunes := []rune(c.Alpha)     // alphabetRunes = ['A','B',...,'Z']
	keyLength := len(KeyRunes)           // keyLength = 8
	alphabetLength := len(alphabetRunes) // alphabetLength = 26

	result := make([]rune, len(str)) // result = []

	for i, char := range str { // i=0, char='H'; i=1, char='E'; i=2, char='L'; i=3, char='L'; i=4, char='O'
		alphaIndex := strings.IndexRune(c.Alpha, char)  // alphaIndex for 'H' = 7
		keyChar := KeyRunes[i%keyLength]                // keyChar for i=0:'P', i=1:'A', i=2:'S', i=3:'S', i=4:'W', i=5:'O', i=6:'R', i=7:'D', i=8:'P'...
		keyIndex := strings.IndexRune(c.Alpha, keyChar) // keyIndex for 'P' = 15 (from the alphabet runes, P is in 15th position)

		if alphaIndex == -1 || keyIndex == -1 {
			result = append(result, char) // Non-alphabet characters remain unchanged
			continue
		}
		newIndex := (alphaIndex + keyIndex) % alphabetLength // For 'H' + 'P': (7+15)%26 = 22 (we use the 7 from 'H' and the 15 from 'P')
		result = append(result, alphabetRunes[newIndex])     // result = ['W'] after first loop ('W' is at position 22 in alphabet)
	}
	// For "HELLO" with "PASSWORD":
	// i=0: 'H' + 'P' -> 'W'
	// i=1: 'E' + 'A' -> 'E'
	// i=2: 'L' + 'S' -> (11+18)%26=3 -> 'D'
	// i=3: 'L' + 'S' -> (11+18)%26=3 -> 'D'
	// i=4: 'O' + 'W' -> (14+22)%26=10 -> 'K'

	return string(result) // returns "WEDDK"
}

func (c VigenèreCipher) Decode(str string) string {
	// c.Key = "PASSWORD" (same as before)
	// c.Alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" (same as before)
	// str = "WEDDK" (the encoded message)

	KeyRunes := []rune(c.Key)        // KeyRunes = ['P','A','S','S','W','O','R','D']
	alphabetRunes := []rune(c.Alpha) // alphabetRunes = ['A','B',...,'Z']

	keyLength := len(KeyRunes)           // keyLength = 8
	alphabetLength := len(alphabetRunes) // alphabetLength = 26

	result := make([]rune, 0, len(str)) // result = []

	for i, char := range str { // i=0, char='W'; i=1, char='E'; i=2, char='D'; i=3, char='D'; i=4, char='K'
		alphaIndex := strings.IndexRune(c.Alpha, char)  // alphaIndex for 'W' = 22
		keyChar := KeyRunes[i%keyLength]                // keyChar for i=0:'P', i=1:'A', i=2:'S', i=3:'S', i=4:'W', i=5:'O', i=6:'R', i=7:'D', i=8:'P'...
		keyIndex := strings.IndexRune(c.Alpha, keyChar) // keyIndex for 'P' = 15 (from the alphabet runes, P is in 15th position)

		if alphaIndex == -1 || keyIndex == -1 {
			result = append(result, char) // Non-alphabet characters remain unchanged
			continue
		}
		newIndex := (alphaIndex - keyIndex + alphabetLength) % alphabetLength // For 'W' - 'P': (22-15+26)%26 = 7 (we use the 22 from 'W' and the 15 from 'P')
		result = append(result, alphabetRunes[newIndex])                      // result = ['H'] after first loop ('H' is at position 7 in alphabet)

		// For "WEDDK" with "PASSWORD":
		// i=0: 'W' - 'P' -> 'H'
		// i=1: 'E' - 'A' -> 'E'
		// i=2: 'D' - 'S' -> (3-18+26)%26=11 -> 'L'
		// i=3: 'D' - 'S' -> (3-18+26)%26=11 -> 'L'
		// i=4: 'K' - 'W' -> (10-22+26)%26=14 -> 'O'

	}

	return string(result) // returns "HELLO"

}
