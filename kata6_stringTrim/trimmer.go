package kata6_stringTrim

import "strings"

func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
