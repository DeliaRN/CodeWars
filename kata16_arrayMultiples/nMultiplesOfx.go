package kata16arraymultiples

func CountBy(x, n int) []int {
	if x > 0 && n > 0 {

		result := make([]int, n) // Create a slice of size n
		for i := 1; i <= n; i++ {
			result[i-1] = x * i
		}
		return result
	}
	return []int{}
}

// Example input: x = 3, n = 4
// Creates a slice of size 4: [0, 0, 0, 0]
// i = 1 -> result[0] = 3 * 1 = 3 -> [3, 0, 0, 0]
// i = 2 -> result[1] = 3 * 2 = 6 -> [3, 6, 0, 0]
// i = 3 -> result[2] = 3 * 3 = 9 -> [3, 6, 9, 0]
// i = 4 -> result[3] = 3 * 4 = 12 -> [3, 6, 9, 12]
// Returns: [3, 6, 9, 12]
// If x <= 0 or n <= 0, returns an empty slice: []

func AlternativeCountBy(x, n int) []int {
	arr := []int{}

	for i := 1; i <= n; i++ {
		arr = append(arr, i*x)
	}

	return arr
}
