package lgame

import (
	"testing"
)

func TestRotateCoordinate90Cw(t *testing.T) {
	cases := map[Coordinate]Coordinate{
		{0, 0}: {3, 0},
		{1, 0}: {3, 1},
		{2, 0}: {3, 2},
		{3, 0}: {3, 3},
		{0, 1}: {2, 0},
		{1, 1}: {2, 1},
		{2, 1}: {2, 2},
		{3, 1}: {2, 3},
		{0, 2}: {1, 0},
		{1, 2}: {1, 1},
		{2, 2}: {1, 2},
		{3, 2}: {1, 3},
		{0, 3}: {0, 0},
		{1, 3}: {0, 1},
		{2, 3}: {0, 2},
		{3, 3}: {0, 3},
	}

	for coord, expected := range cases {
		result := rotateCoordinate90Cw(DefaultSettings(), coord)
		if result != expected {
			t.Errorf("Coordinate %v should rotate to %v, but gave %v", coord, expected, result)
		}
	}
}

func TestRotateLPiece90Cw(t *testing.T) {
	cases := map[LPiece]LPiece{
		{{0, 2}, {0, 1}, {0, 0}, {1, 2}}: {{1, 0}, {2, 0}, {3, 0}, {1, 1}},
	}

	for piece, expected := range cases {
		result := rotateLPiece90Cw(DefaultSettings(), piece)
		if result != expected {
			t.Errorf("L piece %v should rotate to %v, but gave %v", piece, expected, result)
		}
	}
}

func TestFlipCoordinateHor(t *testing.T) {
	cases := map[Coordinate]Coordinate{
		{0, 0}: {3, 0},
		{1, 1}: {2, 1},
		{2, 2}: {1, 2},
		{3, 3}: {0, 3},
	}

	for coord, expected := range cases {
		result := flipCoordinateHor(DefaultSettings(), coord)
		if result != expected {
			t.Errorf("Coordinate %v should flip to %v, but gave %v", coord, expected, result)
		}
	}
}
