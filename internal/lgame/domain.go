package lgame

type GameState struct {
	blue     LPiece
	red      LPiece
	neutrals [2]NeutralPiece
}

type Coordinate struct {
	x int
	y int
}

type LPiece [4]Coordinate

type NeutralPiece Coordinate

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
