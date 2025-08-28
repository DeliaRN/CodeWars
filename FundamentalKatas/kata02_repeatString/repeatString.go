package kata02_repeatString

func RepeatString(n int, s string) string {
	result := ""
	if nIsNotNegative(n) {
		for i := 0; i < n; i++ {
			result += s
		}
		return result
	} else {
		return result
	}
}

func nIsNotNegative(n int) bool {
	return n > 0
}
