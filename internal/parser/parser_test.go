package parser

import (
	"infinibattle-l-game/internal/lgame"
	"testing"
)

func TestParseGameState(t *testing.T) {
	state := ParseGameState("{\"gameState\":{\"board\":{\"board\":[[4,0,0,0],[1,2,2,2],[1,1,1,2],[0,0,0,4]]},\"scorePlayer0\":7,\"scorePlayer1\":50},\"turn\":1,\"player\":1}")

	if state.PlayerTurn != lgame.PlayerRed {
		t.Errorf("Wrong turn: %d", state.PlayerTurn)
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

	players := []lgame.Player{
		{
			Piece: lgame.LPiece{{0, 1}, {0, 2}, {1, 2}, {2, 2}},
			Score: 7,
		},
		{
			Piece: lgame.LPiece{{1, 1}, {2, 1}, {3, 1}, {3, 2}},
			Score: 50,
		},
	}
	for i := 0; i < len(players); i++ {
		if state.Players[i] != players[i] {
			t.Errorf("Wrong player data")
		}
	}
}
