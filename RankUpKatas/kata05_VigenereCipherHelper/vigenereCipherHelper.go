package kata05vigenerecipherhelper

import "strings"

type VigenèreCipher struct {
	Key   string
	Alpha string
}

func (c VigenèreCipher) Encode(str string) string {

	KeyRunes := []rune(c.Key)
	alphabetRunes := []rune(c.Alpha)
	keyLength := len(KeyRunes)
	alphabetLength := len(alphabetRunes)

	result := make([]rune, len(str))

	for i, char := range str {
		alphaIndex := strings.IndexRune(c.Alpha, char)
		keyChar := KeyRunes[i%keyLength]
		keyIndex := strings.IndexRune(c.Alpha, keyChar)

		if alphaIndex == -1 || keyIndex == -1 {
			result = append(result, char) // Non-alphabet characters remain unchanged
			continue
		}
		newIndex := (alphaIndex + keyIndex) % alphabetLength
		result = append(result, alphabetRunes[newIndex])
	}

	return string(result)
}

func (c VigenèreCipher) Decode(str string) string {

	KeyRunes := []rune(c.Key)
	alphabetRunes := []rune(c.Alpha)

	keyLength := len(KeyRunes)
	alphabetLength := len(alphabetRunes)

	result := make([]rune, 0, len(str))

	for i, char := range str {
		alphaIndex := strings.IndexRune(c.Alpha, char)
		keyChar := KeyRunes[i%keyLength]
		keyIndex := strings.IndexRune(c.Alpha, keyChar)

		if alphaIndex == -1 || keyIndex == -1 {
			result = append(result, char) // Non-alphabet characters remain unchanged
			continue
		}
		newIndex := (alphaIndex - keyIndex + alphabetLength) % alphabetLength
		result = append(result, alphabetRunes[newIndex])
	}

	return string(result)

}
