package kata14digits

func Digits(num int) []int {

	if num == 0 {
		return []int{0}
	}

	var allDigits []int

	//12345 extract digits -> [1,2,3,4,5]
	for num > 0 {
		allDigits = append([]int{num % 10}, allDigits...) // Prepend to maintain order
		num /= 10
	}

	var result []int
	for i := 0; i < len(allDigits); i++ {
		for j := i + 1; j < len(allDigits); j++ {
			result = append(result, allDigits[i]+allDigits[j])
		}
	}

	return result
}
