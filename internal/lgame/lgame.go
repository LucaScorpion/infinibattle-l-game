package lgame

func GetNextState(cur GameState) GameState {
	// TODO
	return getPossibleNextStates(DefaultSettings(), cur)[0]
}

func getPossibleNextStates(settings GameSettings, cur GameState) []GameState {
	lMoves := getLShapeMoves(settings, cur)

	var totalStates []GameState
	for _, s := range lMoves {
		totalStates = append(totalStates, getNeutralMoves(settings, s)...)
	}

	return totalStates
}

func getLShapeMoves(settings GameSettings, state GameState) []GameState {
	grid := getOccupation(state)
	curPlayerOcc := playerIndexToOccupation[state.Turn]
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
				for i, c := range lShape {
					newC := Coordinate{c.X + offsetX, c.Y + offsetY}

					// Check if the piece is in bounds.
					if newC.X < 0 || newC.Y < 0 || newC.X >= settings.BoardWidth || newC.Y >= settings.BoardHeight {
						continue ROWS
					}

					newPlacement[i] = newC
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

				// Create and append the new state.
				newState := state
				newState.Turn = (state.Turn + 1) % 2
				newState.Players[state.Turn] = newPlacement
				nextStates = append(nextStates, newState)
			}
		}
	}

	return nextStates
}

func getNeutralMoves(settings GameSettings, state GameState) []GameState {
	var nextStates []GameState

	for i, n := range state.Neutrals {
		grid := getOccupation(state)
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

func getOccupation(state GameState) occupationGrid {
	occupied := occupationGrid{}

	// Add the L pieces.
	for i, p := range state.Players {
		for _, c := range p {
			occupied[c] = playerIndexToOccupation[PlayerIndex(i)]
		}
	}

	// Add the neutral pieces.
	for _, n := range state.Neutrals {
		occupied[Coordinate{n.X, n.Y}] = occupiedNeutral
	}

	return occupied
}

func isSamePlacement(piece LPiece, grid occupationGrid) bool {
	for _, c := range piece {
		// Here we assume that the L piece placement is valid,
		// i.e. it only ever overlaps with itself.
		if _, ok := grid[c]; !ok {
			return false
		}
	}
	return true
}
