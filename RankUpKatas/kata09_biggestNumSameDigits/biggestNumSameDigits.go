package kata09biggestnumsamedigits

func NextBigger(n int) int {

	var digits []int
	// Extract digits from the number
	for n > 0 {
		digits = append([]int{n % 10}, digits...)
		n /= 10
	}
	length := len(digits)

	orderedDigits := make([]int, length)
	copy(orderedDigits, digits)

	// Find the pivot point
	i := length - 2
	for i >= 0 && digits[i] >= digits[i+1] {
		i--
	}

	// If no pivot point is found, return -1
	if i < 0 {
		return -1
	}

	j := length - 1
	for digits[j] <= digits[i] {
		j--
	}

	// Swap the pivot with the rightmost successor
	digits[i], digits[j] = digits[j], digits[i]

	// Reverse digits to the right of the pivot
	for left, right := i+1, length-1; left < right; left, right = left+1, right-1 {
		digits[left], digits[right] = digits[right], digits[left]
	}
	// Convert digits back to a number
	result := 0
	for _, digit := range digits {
		result = result*10 + digit
	}
	return result

}
