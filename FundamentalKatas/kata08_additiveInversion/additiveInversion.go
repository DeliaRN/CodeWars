package kata08_additiveInversion

/*
Given a set of numbers, return the additive inverse of each.
Each positive becomes negatives, and the negatives become positives.

[1, 2, 3, 4, 5] --> [-1, -2, -3, -4, -5]
[1, -2, 3, -4, 5] --> [-1, 2, -3, 4, -5]
[] --> []
You can assume that all values are integers. Do not mutate the input array.
*/

func AdditiveInversion(arr []int) []int {
	result := make([]int, len(arr))
	for i, num := range arr {
		result[i] = -num
	}
	return result
}
