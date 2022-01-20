package lgame

func GetPossibleNextStates(settings GameSettings, cur GameState) []GameState {
	lMoves := GetLShapeMoves(settings, cur)

	var totalStates []GameState
	for _, s := range lMoves {
		totalStates = append(totalStates, GetNeutralMoves(settings, s)...)
	}

	return totalStates
}

func GetLShapeMoves(settings GameSettings, state GameState) []GameState {
	grid := GetOccupation(state)
	curPlayerOcc := playerIndexToOccupation[state.PlayerTurn]
	var nextStates []GameState

	for _, lShape := range lShapes {
		lCorner := lShape[0]

		for x := 0; x < settings.BoardWidth; x++ {
		ROWS:
			for y := 0; y < settings.BoardHeight; y++ {
				// Get the L piece offset.
				offsetX := x - lCorner.X
				offsetY := y - lCorner.Y

				// Get the new L piece placement.
				newPlacement := LPiece{}
				isCornerPoint := false
				for i, c := range lShape {
					newC := Coordinate{c.X + offsetX, c.Y + offsetY}

					// Check if the piece is in bounds.
					if newC.X < 0 || newC.Y < 0 || newC.X >= settings.BoardWidth || newC.Y >= settings.BoardHeight {
						continue ROWS
					}

					newPlacement[i] = newC

					// Check if the piece scores a corner point.
					if (newC.X == 0 || newC.X == settings.BoardWidth-1) && (newC.Y == 0 || newC.Y == settings.BoardHeight-1) {
						isCornerPoint = true
					}
				}

				// Check if any of the placement coordinates are occupied by something other than the current player.
				for _, c := range newPlacement {
					if o, ok := grid[c]; ok && o != curPlayerOcc {
						continue ROWS
					}
				}

				// Check if the placement is the same as the previous placement.
				if isSamePlacement(newPlacement, grid) {
					continue
				}

				// Create the new state.
				newState := state
				newState.Players[state.PlayerTurn].Piece = newPlacement
				if isCornerPoint {
					newState.Players[state.PlayerTurn].Score++
				}
				newState.PlayerTurn = OtherPlayer(state.PlayerTurn)
				nextStates = append(nextStates, newState)
			}
		}
	}

	return nextStates
}

func GetNeutralMoves(settings GameSettings, state GameState) []GameState {
	var nextStates []GameState

	for i, n := range state.Neutrals {
		grid := GetOccupation(state)
		delete(grid, Coordinate(n))

		for x := 0; x < settings.BoardWidth; x++ {
			for y := 0; y < settings.BoardHeight; y++ {
				check := Coordinate{x, y}

				// Check if the space is occupied.
				if _, ok := grid[check]; ok {
					continue
				}

				// Create and append the new state.
				newState := state
				newState.Neutrals[i] = NeutralPiece(check)
				nextStates = append(nextStates, newState)
			}
		}
	}

	return nextStates
}

func OtherPlayer(p PlayerIndex) PlayerIndex {
	return (p + 1) % 2
}

func GetOccupation(state GameState) OccupationGrid {
	occupied := OccupationGrid{}

	// Add the L pieces.
	for i, p := range state.Players {
		for _, c := range p.Piece {
			occupied[c] = playerIndexToOccupation[PlayerIndex(i)]
		}
	}

	// Add the neutral pieces.
	for _, n := range state.Neutrals {
		occupied[Coordinate(n)] = occupiedNeutral
	}

	return occupied
}

func isSamePlacement(piece LPiece, grid OccupationGrid) bool {
	for _, c := range piece {
		// Here we assume that the L piece placement is valid,
		// i.e. it only ever overlaps with itself.
		if _, ok := grid[c]; !ok {
			return false
		}
	}
	return true
}
