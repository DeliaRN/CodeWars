package kata10_rockpaperscissors

func newGame(playerOne, playerTwo string) string {

	if playerOne == playerTwo {
		return "Draw!"
	}

	switch playerOne {
	case "rock":
		if playerTwo == "scissors" {
			return "Player 1 wins!"
		}
		return "Player 2 wins!"
	case "paper":
		if playerTwo == "rock" {
			return "Player 1 wins!"
		}
		return "Player 2 wins!"
	case "scissors":
		if playerTwo == "paper" {
			return "Player 1 wins!"
		}
		return "Player 2 wins!"
	default:
		return "Invalid game!"
	}

}
