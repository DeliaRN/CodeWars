package kata9_stringtonumb

import (
	"strconv"
)

func stringConversor(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return num
}
