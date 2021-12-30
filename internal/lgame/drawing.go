package lgame

import (
	"fmt"
	"strings"
)

func drawState(settings GameSettings, state GameState) string {
	grid := make([][]string, settings.BoardHeight)

	// Initialize the empty grid.
	for y := 0; y < settings.BoardHeight; y++ {
		grid[y] = make([]string, settings.BoardWidth)
		for x := 0; x < settings.BoardHeight; x++ {
			grid[y][x] = "\x1b[37;47m  \x1b[0m"
		}
	}

	// Add the neutral pieces.
	for _, n := range state.Neutrals {
		grid[n.Y][n.X] = "\x1b[30;40mNN\x1b[0m"
	}

	// Add the players.
	for i, p := range state.Players {
		color := 1
		if i == int(PlayerBlue) {
			color = 4
		}

		for _, c := range p {
			grid[c.Y][c.X] = fmt.Sprintf("\x1b[3%d;4%dm%d%d\x1b[0m", color, color, i, i)
		}
	}

	lines := make([]string, len(grid))
	for i, l := range grid {
		lines[i] = strings.Join(l, "")
	}
	return "\n" + strings.Join(lines, "\n")
}
