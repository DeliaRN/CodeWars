package kata06josephussurvivor

func JosephusSurvivor(n, k int) int {

	step := k - 1
	result := make([]int, n)

	// If n = 7, k = 3, it means there are 7 people and every 3rd person will be eliminated
	// The step is k - 1 because we count the current person as well
	// Initialize the result slice with people numbered from 1 to n
	for i := 0; i < n; i++ {
		result[i] = i + 1
	} //This loop fills the slice with numbers from 1 to n

	// Start eliminating people until only one person is left
	index := 0
	for len(result) > 1 {
		index = (index + step) % len(result) // Calculate the index of the person to be eliminated
		result = append(result[:index], result[index+1:]...)
		//This creates a new slice that skips the element at 'index', effectively removing it
		//[:index] gives all elements before 'index'
		//[index+1:] gives all elements after 'index'
		//Appending these two slices together results in a new slice without the eliminated person

	}
	return result[0] // Return the last remaining person

	// Example: n = 7, k = 3
	// Initial result: [1, 2, 3, 4, 5, 6, 7]
	// Step = 2 (k - 1)
	// Iteration 1: index = (0 + 2) % 7 = 2 -> Remove 3 -> [1, 2, 4, 5, 6, 7]
	// Iteration 2: index = (2 + 2) % 6 = 4 -> Remove 6 -> [1, 2, 4, 5, 7]
	// Iteration 3: index = (4 + 2) % 5 = 1 -> Remove 2 -> [1, 4, 5, 7]
	// Iteration 4: index = (1 + 2) % 4 = 3 -> Remove 7 -> [1, 4, 5]
	// Iteration 5: index = (3 + 2) % 3 = 2 -> Remove 5 -> [1, 4]
	// Iteration 6: index = (2 + 2) % 2 = 0 -> Remove 1 -> [4]
	// Last remaining person is 4
}
