package bot

import (
	"infinibattle-l-game/internal/lgame"
	"math"
	"testing"
	"time"
)

/*
A more difficult state.
Blue's turn.
L-shape moves: 17.
Neutral moves: 14.
Total possible moves: 238.

  x 0 1 2 3
y ┌─────────┐
0 │ R R N □ │
1 │ R □ □ □ │
2 │ R B B B │
3 │ N □ □ B │
  └─────────┘
*/
func getDifficultState() lgame.GameState {
	return lgame.GameState{
		PlayerTurn: lgame.PlayerBlue,
		Players: [2]lgame.Player{
			{
				Piece: lgame.LPiece{
					{0, 2},
					{0, 1},
					{0, 0},
					{1, 0},
				},
				Score: 0,
			},
			{
				Piece: lgame.LPiece{
					{1, 2},
					{2, 2},
					{3, 2},
					{3, 3},
				},
				Score: 0,
			},
		},
		Neutrals: [2]lgame.NeutralPiece{
			{2, 0},
			{0, 3},
		},
	}
}

func BenchmarkGetNextState(b *testing.B) {
	totalTime := 0.0
	worstTime := 0.0

	for i := 0; i < b.N; i++ {
		startTime := time.Now()

		getNextState(lgame.DefaultSettings(), getDifficultState())

		thinkTime := time.Now().Sub(startTime).Seconds()
		worstTime = math.Max(worstTime, thinkTime)
		totalTime = totalTime + thinkTime
	}

	if b.N > 1 {
		b.Logf("Average time: %.3f seconds", totalTime/float64(b.N))
		b.Logf("Worst time: %.3f seconds", worstTime)
	}
}
