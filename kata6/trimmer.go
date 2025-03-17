package kata6

import "strings"

func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
