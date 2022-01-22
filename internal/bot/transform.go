package bot

import (
	"infinibattle-l-game/internal/lgame"
)

func rotate90Cw(settings lgame.GameSettings, state lgame.GameState) lgame.GameState {
	rotated := state

	for i, player := range state.Players {
		rotated.Players[i].Piece = rotateLPiece90Cw(settings, player.Piece)
	}

	for i, neutral := range state.Neutrals {
		rotated.Neutrals[i] = lgame.NeutralPiece(rotateCoordinate90Cw(settings, lgame.Coordinate(neutral)))
	}

	return rotated
}

func rotateCoordinate90Cw(settings lgame.GameSettings, coord lgame.Coordinate) lgame.Coordinate {
	return lgame.Coordinate{
		X: -(coord.Y - (settings.BoardHeight - 1)),
		Y: coord.X,
	}
}

func rotateLPiece90Cw(settings lgame.GameSettings, piece lgame.LPiece) lgame.LPiece {
	rotated := lgame.LPiece{}
	for i, coord := range piece {
		rotated[i] = rotateCoordinate90Cw(settings, coord)
	}
	return rotated
}
