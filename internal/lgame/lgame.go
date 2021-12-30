package lgame

func getPossibleNextStates(settings GameSettings, cur GameState, playerTurn int) []GameState {
	// TODO: move neutral pieces.
	return getLShapeMoves(settings, cur, playerTurn)
}

func getLShapeMoves(settings GameSettings, cur GameState, playerTurn int) []GameState {
	occupation := getOccupation(settings, cur, playerTurn)
	playerOccupation := getPlayerOccupation(cur.Players[playerTurn])
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

				// Check if any of the placement coordinates are occupied.
				for _, c := range newPlacement {
					if occupation[c] {
						continue ROWS
					}
				}

				// Check if the placement is the same as the previous placement.
				if isSamePlacement(newPlacement, playerOccupation) {
					continue
				}

				// Create and append the new state.
				newState := cur
				newState.Players[playerTurn] = newPlacement
				nextStates = append(nextStates, newState)
			}
		}
	}

	return nextStates
}

func getOccupation(settings GameSettings, state GameState, playerTurn int) map[Coordinate]bool {
	occupied := map[Coordinate]bool{}

	// Initialize all occupied coords to false.
	for x := 0; x < settings.BoardWidth; x++ {
		for y := 0; y < settings.BoardHeight; y++ {
			occupied[Coordinate{x, y}] = false
		}
	}

	// Add the L pieces.
	for i, p := range state.Players {
		// Skip the current player.
		if i == playerTurn {
			continue
		}

		for _, c := range p {
			occupied[c] = true
		}
	}

	// Add the neutral pieces.
	for _, n := range state.Neutrals {
		occupied[Coordinate{n.X, n.Y}] = true
	}

	return occupied
}

func getPlayerOccupation(piece LPiece) map[Coordinate]bool {
	occupation := map[Coordinate]bool{}

	for _, c := range piece {
		occupation[c] = true
	}

	return occupation
}

func isSamePlacement(piece LPiece, playerOccupation map[Coordinate]bool) bool {
	for _, c := range piece {
		if _, ok := playerOccupation[c]; !ok {
			return false
		}
	}
	return true
}
