package kata1_smallestInt

func SmallestIntegerFinder(numbers []int) int {
	smallest := numbers[0]
	for _, number := range numbers {
		if number < smallest {
			smallest = number
		}
	}
	return smallest
}
