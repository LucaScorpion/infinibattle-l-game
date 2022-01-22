package lgame

func AllTransforms(settings GameSettings, state GameState) []GameState {
	transforms := allRotations(settings, state)
	transforms = append(transforms, allRotations(settings, flipHor(settings, state))...)
	return transforms
}

func allRotations(settings GameSettings, state GameState) []GameState {
	rotations := make([]GameState, 4)
	rotations[0] = state

	for i := 0; i < 3; i++ {
		rotations[i+1] = rotate90Cw(settings, state)
	}

	return rotations
}

func rotate90Cw(settings GameSettings, state GameState) GameState {
	rotated := state

	for i, player := range state.Players {
		rotated.Players[i].Piece = rotateLPiece90Cw(settings, player.Piece)
	}

	for i, neutral := range state.Neutrals {
		rotated.Neutrals[i] = NeutralPiece(rotateCoordinate90Cw(settings, Coordinate(neutral)))
	}

	return rotated
}

func rotateCoordinate90Cw(settings GameSettings, coord Coordinate) Coordinate {
	return Coordinate{
		X: -(coord.Y - (settings.BoardHeight - 1)),
		Y: coord.X,
	}
}

func rotateLPiece90Cw(settings GameSettings, piece LPiece) LPiece {
	rotated := LPiece{}
	for i, coord := range piece {
		rotated[i] = rotateCoordinate90Cw(settings, coord)
	}
	return rotated
}

func flipHor(settings GameSettings, state GameState) GameState {
	flipped := state

	for i, player := range state.Players {
		flipped.Players[i].Piece = flipLPieceHor(settings, player.Piece)
	}

	for i, neutral := range state.Neutrals {
		flipped.Neutrals[i] = NeutralPiece(flipCoordinateHor(settings, Coordinate(neutral)))
	}

	return flipped
}

func flipCoordinateHor(settings GameSettings, coord Coordinate) Coordinate {
	return Coordinate{
		X: -(coord.X - (settings.BoardWidth - 1)),
		Y: coord.Y,
	}
}

func flipLPieceHor(settings GameSettings, piece LPiece) LPiece {
	flipped := LPiece{}
	for i, coord := range piece {
		flipped[i] = flipCoordinateHor(settings, coord)
	}
	return flipped
}
