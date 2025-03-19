package kata04_descInts

/*
DESCRIPTION:
Build a function that returns an array of integers from n to 1 where n>0.
Example : n=5 --> [5,4,3,2,1]
*/

func descendingArray(number int) []int {
	array := make([]int, number)
	var temp = number

	for i := 0; i < number; i++ {
		array[i] = temp
		temp--
	}
	return array
}
