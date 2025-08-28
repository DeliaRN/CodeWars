package kata06_stringTrim

import "strings"

func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
