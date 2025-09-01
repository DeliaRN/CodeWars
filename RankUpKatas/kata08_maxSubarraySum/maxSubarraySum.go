package kata08maxsubarraysum

func MaximumSubarraySum(numbers []int) int {

	if len(numbers) == 0 {
		return 0
	}

	var maxSum int = 0
	var currentSum int = 0

	for _, num := range numbers {

		currentSum = max(currentSum+num, num) // Choose to add or start new subarray
		maxSum = max(maxSum, currentSum)      // Update maxSum if currentSum is greater
	}

	if maxSum < 0 {
		return 0
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
