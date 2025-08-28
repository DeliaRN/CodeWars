package kata25upperchecker

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	//first convert s to a regular string
	// then use strings.ToUpper to convert it to uppercase
	// and then compares it with the original string
	return string(s) == strings.ToUpper(string(s))
}
