package kata19kittenpup

// SHORTER WAY (as I wanted as first)
func ShorterCalculateYears(years int) (result [3]int) {
	switch years {
	case 1:
		result = [3]int{1, 15, 15}
	case 2:
		result = [3]int{2, 24, 24}
	default:
		result = [3]int{years, 24 + 4*(years-2), 24 + 5*(years-2)}
	}
	return
}

// My second approach, with if statements
func CalculateYears(years int) (result [3]int) {
	result[0] = years
	if years <= 0 {
		return result
	}

	if years == 1 {
		result[1] = 15
		result[2] = 15
		return result
	} else if years == 2 {
		result[1] = 24
		result[2] = 24
		return result
	} else {
		remaining := years - 2
		result[1] = remaining*4 + 24
		result[2] = remaining*5 + 24
		return result
	}
}

// ALTERNATIVE WAY WITH SWITCH-CASE
func ALTCalculateYears(years int) (result [3]int) {

	result[0] = years
	switch {
	case years <= 0:
		return result
	case years == 1:
		result[1] = 15
		result[2] = 15
		return result
	case years == 2:
		result[1] = 24
		result[2] = 24
		return result
	default:
		remaining := years - 2
		result[1] = remaining*4 + 24
		result[2] = remaining*5 + 24
		return result
	}
}
