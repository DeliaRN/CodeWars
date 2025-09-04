package kata10snakes

import (
	"fmt"
)

// SnakesLadders - Stucture that game is played on
type SnakesLadders struct {
	player1       int
	player2       int
	currentPlayer int
	gameOver      bool
}

// NewGame - Method that starts a new game for the SnakesLadders struct
func (sl *SnakesLadders) NewGame() {

	sl.player1 = 0
	sl.player2 = 0
	sl.currentPlayer = 1
	sl.gameOver = false
}

// Play - Method that makes turn givem a doce roll from die1 and die2 for the SnakesLadders struct
// Player that dice is for should automatically be determined based on rules
func (sl *SnakesLadders) Play(die1, die2 int) string {

	if sl.gameOver {
		return "Game over!"
	}

	activePlayer := sl.currentPlayer

	// Select current player's position
	var position *int
	if activePlayer == 1 {
		position = &sl.player1
	} else {
		position = &sl.player2
	}

	// Move player
	move := die1 + die2
	*position += move

	// Bounce back if over 100
	if *position > 100 {
		*position = 100 - (*position - 100)
	}

	// Check for snakes or ladders
	*position = checkPosition(*position)

	// Check for win
	if *position == 100 {
		sl.gameOver = true
		return fmt.Sprintf("Player %d Wins!", activePlayer)
	}

	// Build message before switching players
	msg := fmt.Sprintf("Player %d is on square %d", activePlayer, *position)

	// Change turn if no double
	if die1 != die2 {
		if sl.currentPlayer == 1 {
			sl.currentPlayer = 2
		} else {
			sl.currentPlayer = 1
		}
	}
	return msg
}

func checkPosition(pos int) int {
	switch pos {
	case 2:
		return 38
	case 7:
		return 14
	case 8:
		return 31
	case 15:
		return 26
	case 16:
		return 6
	case 21:
		return 42
	case 28:
		return 84
	case 36:
		return 44
	case 46:
		return 25
	case 49:
		return 11
	case 51:
		return 67
	case 62:
		return 19
	case 64:
		return 60
	case 71:
		return 91
	case 74:
		return 53
	case 78:
		return 98
	case 87:
		return 94
	case 89:
		return 68
	case 92:
		return 88
	case 95:
		return 75
	case 99:
		return 80
	default:
		return pos
	}
}
