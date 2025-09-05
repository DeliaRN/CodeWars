package kata12_greedisgood

func Score(dice [5]int) int {
	score := 0
	//if there is any 0 in the dice, return -1
	for _, die := range dice {
		if die == 0 {
			return -1
		}
	}

	counts := make(map[int]int)
	for _, die := range dice {
		counts[die]++
	}

	for die, count := range counts {
		if count >= 3 {
			if die == 1 {
				score += 1000
			} else {
				score += die * 100
			}
			count -= 3
		}
		if die == 1 {
			score += count * 100
		} else if die == 5 {
			score += count * 50
		}
	}

	return score
}
