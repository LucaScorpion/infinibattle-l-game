package lgame

type GameState struct {
	Players  [2]LPiece
	Neutrals [2]NeutralPiece
}

type Coordinate struct {
	X int
	Y int
}

type LPiece [4]Coordinate

type NeutralPiece Coordinate

type PlayerIndex int

const (
	PlayerRed PlayerIndex = iota
	PlayerBlue
)

type occupation int

const (
	occupiedRed occupation = iota
	occupiedBlue
	occupiedNeutral
)

type occupationGrid map[Coordinate]occupation

var playerIndexToOccupation = map[PlayerIndex]occupation{
	PlayerRed:  occupiedRed,
	PlayerBlue: occupiedBlue,
}

/*
Order of coordinates of the L-shapes:

  x 0 1 2 3
y ┌─────────┐
0 │ 2 □ □ □ │
1 │ 1 □ □ □ │
2 │ 0 3 □ □ │
3 │ □ □ □ □ │
  └─────────┘

The corner piece is always the first coordinate in the list.
*/
var lShapes = []LPiece{
	// Long side vertical.
	{{0, 2}, {0, 1}, {0, 0}, {1, 2}}, // L
	{{1, 2}, {1, 1}, {1, 0}, {0, 2}}, // ⅃
	{{0, 0}, {0, 1}, {0, 2}, {1, 0}}, // Γ
	{{1, 0}, {1, 1}, {1, 2}, {0, 0}}, // ꓶ
	// Long side horizontal.
	{{0, 1}, {1, 1}, {2, 1}, {0, 0}}, // L
	{{2, 1}, {1, 1}, {0, 1}, {2, 0}}, // ⅃
	{{0, 0}, {1, 0}, {2, 0}, {0, 1}}, // Γ
	{{2, 0}, {1, 0}, {0, 0}, {2, 1}}, // ꓶ
}
