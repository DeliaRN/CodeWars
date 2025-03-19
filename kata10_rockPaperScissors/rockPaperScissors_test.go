package kata10_rockpaperscissors

import "testing"

func TestNewGame(t *testing.T) {

	/* ONE WAY OF DOING IT, TOO MUCH DUPLICATED CODE
	t.Run("Game 1", func(t *testing.T) {
		got := newGame("rock", "scissors")
		want := "Player 1 won!"
		assertWinner(t, got, want)
	}) */

	tests := []struct {
		playerOne string
		playerTwo string
		expected  string
	}{
		{"rock", "scissors", "Player 1 wins!"},
		{"rock", "rock", "Draw!"},
		{"rock", "paper", "Player 2 wins!"},
		{"paper", "scissors", "Player 2 wins!"},
		{"paper", "rock", "Player 1 wins!"},
		{"paper", "paper", "Draw!"},
		{"scissors", "scissors", "Draw!"},
		{"scissors", "rock", "Player 2 wins!"},
		{"scissors", "paper", "Player 1 wins!"},
	}

	for _, test := range tests {
		result := newGame(test.playerOne, test.playerTwo)
		if result != test.expected {
			t.Errorf("The game %q vs %q should result in %q, but got %q", test.playerOne, test.playerTwo, test.expected, result)
		}
	}
}
