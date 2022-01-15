package parser

import (
	"encoding/json"
	"fmt"
	"infinibattle-l-game/internal/lgame"
)

func ParseGameState(in string) lgame.GameState {
	turn := parseTurnState(in)
	board := turn.GameState.Board.Board

	playerOne := lgame.LPiece{}
	playerOneLen := 0
	playerTwo := lgame.LPiece{}
	playerTwoLen := 0
	var neutrals []lgame.NeutralPiece

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[y]); x++ {
			switch board[y][x] {
			case empty:
				// Do nothing.
			case player1:
				playerOne[playerOneLen] = lgame.Coordinate{X: x, Y: y}
				playerOneLen++
			case player2:
				playerTwo[playerTwoLen] = lgame.Coordinate{X: x, Y: y}
				playerTwoLen++
			case neutral:
				neutrals = append(neutrals, lgame.NeutralPiece{X: x, Y: y})
			default:
				panic(fmt.Sprintf("Unknown piece type: %d", board[y][x]))
			}
		}
	}

	return lgame.GameState{
		Turn:     lgame.PlayerIndex(turn.Turn - 1),
		Players:  []lgame.LPiece{playerOne, playerTwo},
		Neutrals: neutrals,
	}
}

func parseTurnState(in string) turnState {
	var state turnState
	if err := json.Unmarshal([]byte(in), &state); err != nil {
		panic(err)
	}
	return state
}
