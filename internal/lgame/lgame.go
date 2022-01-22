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
	_, grid := GetOccupation(state)
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
		_, grid := GetOccupation(state)
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

func GetOccupation(state GameState) (bool, OccupationGrid) {
	occupied := OccupationGrid{}

	// Add the L pieces.
	for i, p := range state.Players {
		for _, c := range p.Piece {
			// Check if the square is already occupied.
			if _, ok := occupied[c]; ok {
				return false, occupied
			}

			occupied[c] = playerIndexToOccupation[PlayerIndex(i)]
		}
	}

	// Add the neutral pieces.
	for _, n := range state.Neutrals {
		// Check if the square is already occupied.
		if _, ok := occupied[Coordinate(n)]; ok {
			return false, occupied
		}

		occupied[Coordinate(n)] = occupiedNeutral
	}

	return true, occupied
}

func IsValidMove(from GameState, to GameState) bool {
	playerPiece := to.Players[from.PlayerTurn].Piece
	_, fromOcc := GetOccupation(from)
	toValid, toOcc := GetOccupation(to)
	curPlayerOcc := playerIndexToOccupation[from.PlayerTurn]

	if !toValid {
		return false
	}

	// Check if the L piece doesn't overlap anything other than itself in the old state.
	// We need to check this against the old state because we need to make sure that the
	// L shape moved before any neutral pieces were moved.
	for _, c := range playerPiece {
		if o, ok := fromOcc[c]; ok && o != curPlayerOcc {
			return false
		}
	}

	// Check if the L piece doesn't overlap anything other than itself in the new state.
	for _, c := range playerPiece {
		if o, ok := toOcc[c]; ok && o != curPlayerOcc {
			return false
		}
	}

	// Check if the L piece changed position.
	if isSamePlacement(playerPiece, fromOcc) {
		return false
	}

	// Check if at most one neutral piece moved.
	moved := make(map[NeutralPiece]bool)
	for _, n := range from.Neutrals {
		moved[n] = true
	}
	for _, n := range to.Neutrals {
		delete(moved, n)
	}
	if len(moved) > 1 {
		return false
	}

	return true
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
