package kata14_stringToArray

/*
Write a function to split a string and convert it into an array of words.
Examples (Input ==> Output):
"Robin Singh" ==> ["Robin", "Singh"]
"I love arrays they are my favorite"
==> ["I", "love", "arrays", "they", "are", "my", "favorite"]
*/

import "strings"

func StringToArray(str string) []string {
	var result []string
	result = strings.Split(str, " ")
	return result
}
