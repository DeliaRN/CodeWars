package kata3

/*
You were camping with your friends far away from home, but when it's time to go back, you realize that
your fuel is running out and the nearest pump is 50 miles away! You know that on average,
your car runs on about 25 miles per gallon. There are 2 gallons left.

Write a function that tells you if it is possible to get to the pump or not.
*/

func isFuelEnough(distanceToPump, performance, gallonsLeft int) bool {
	return gallonsLeft*performance >= distanceToPump
}
