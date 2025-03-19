package kata5_hopper

/*
Write a program that finds the summation of every number from 1 to num.
The number will always be a positive integer greater than 0. Your function
only needs to return the result, what is shown between parentheses in the
example below is how you reach that result and it's not part of it,
see the sample tests.
2 -> 3 (1 + 2)
8 -> 36 (1 + 2 + 3 + 4 + 5 + 6 + 7 + 8)
*/

func hopper(input int) int {
	result := 0
	for i := 0; i < (input + 1); i++ {
		result += i
	}
	return result
}
