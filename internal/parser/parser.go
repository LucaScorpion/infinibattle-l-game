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
	var neutrals [2]lgame.NeutralPiece
	neutralsLen := 0

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
				neutrals[neutralsLen] = lgame.NeutralPiece{X: x, Y: y}
				neutralsLen++
			default:
				panic(fmt.Sprintf("Unknown piece type: %d", board[y][x]))
			}
		}
	}

	return lgame.GameState{
		PlayerTurn: lgame.PlayerIndex(turn.Player),
		Players: [2]lgame.Player{
			{
				Piece: playerOne,
				Score: turn.GameState.ScorePlayer0,
			},
			{
				Piece: playerTwo,
				Score: turn.GameState.ScorePlayer1,
			},
		},
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

func GetMoveOutput(nextState lgame.GameState, playerTurn lgame.PlayerIndex) string {
	var place placePiecesCommand
	for i := 0; i < len(nextState.Neutrals); i++ {
		place.NeutralPieceCoordinates[i] = coordinateOutput(lgame.Coordinate(nextState.Neutrals[i]))
	}

	playerPiece := nextState.Players[playerTurn].Piece
	for i := 0; i < len(playerPiece); i++ {
		place.PlayerLPieceCoordinates[i] = coordinateOutput(playerPiece[i])
	}

	bytes, err := json.Marshal(place)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func coordinateOutput(coord lgame.Coordinate) coordinate {
	return coordinate{coord.X, coord.Y}
}
