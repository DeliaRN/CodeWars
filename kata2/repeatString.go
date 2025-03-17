package kata2

func RepeatString(n int, s string) string {
	result := ""
	for i := 0; i < n; i++ {
		result += s
	}

	return result
}
