package kata11_initialscreator

import "strings"

func initialsCreator(fullname string) string {
	// split the name in two:
	parts := strings.Split(fullname, " ")
	// take only the first letter of each word
	initials := strings.ToUpper(string(parts[0][0])) + "." + strings.ToUpper(string(parts[1][0]))

	return initials
}
