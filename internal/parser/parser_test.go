package parser

import (
	"infinibattle-l-game/internal/lgame"
	"testing"
)

func TestParseGameState(t *testing.T) {
	state := ParseGameState("{\"gameState\":{\"board\":{\"board\":[[4,0,0,0],[1,2,2,2],[1,1,1,2],[0,0,0,4]]},\"scorePlayer0\":0,\"scorePlayer1\":0},\"turn\":1,\"player\":1}")

	if state.Turn != lgame.PlayerRed {
		t.Errorf("Wrong turn: %d", state.Turn)
	}
	if len(state.Neutrals) != 2 {
		t.Errorf("Incorrect number of neutral pieces: %d", len(state.Neutrals))
	}
	if len(state.Players) != 2 {
		t.Errorf("Incorrect number of players: %d", len(state.Players))
	}

	neutrals := []lgame.NeutralPiece{
		{0, 0},
		{3, 3},
	}
	for i := 0; i < len(neutrals); i++ {
		if state.Neutrals[i] != neutrals[i] {
			t.Errorf("Wrong position for neutral pieces")
		}
	}

	lPieces := []lgame.LPiece{
		{{0, 1}, {0, 2}, {1, 2}, {2, 2}},
		{{1, 1}, {2, 1}, {3, 1}, {3, 2}},
	}
	for i := 0; i < len(lPieces); i++ {
		if state.Players[i] != lPieces[i] {
			t.Errorf("Wrong position for player pieces")
		}
	}
}
