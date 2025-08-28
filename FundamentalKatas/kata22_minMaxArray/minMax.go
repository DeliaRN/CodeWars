package kata22minmaxarray

// return both the minimum and maximum number of the given list/array.
func MinMax(arr []int) [2]int {
	if len(arr) == 0 {
		return [2]int{}
	} else if len(arr) == 1 {
		return [2]int{arr[0], arr[0]}
	} else {
		min, max := arr[0], arr[0]
		for _, num := range arr {
			if num < min {
				min = num
			}
			if num > max {
				max = num
			}
		}
		return [2]int{min, max}
	}
}
