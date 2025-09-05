package kata13sumintervals

import "sort"

func SumOfIntervals(intervals [][2]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	sum := 0
	low := intervals[0][0]

	for _, value := range intervals {
		if value[1] > low {
			sum += value[1]
			if value[0] > low {
				sum -= value[0]
			} else {
				sum -= low
			}
			low = value[1]
		}
	}

	return sum
}

func WRONGSumOfIntervals(intervals [][2]int) int {
	sum := 0

	for _, interval := range intervals {
		sum += interval[1] - interval[0]
	}

	return sum
}

// WRONG: does not account for overlapping intervals
// e.g. [[1,4],[3,5]] when merged -> [1,5] should return 4
// but this function returns 5 (3 + 2)
