package lgame

type GameSettings struct {
	BoardWidth  int
	BoardHeight int
}

func DefaultSettings() GameSettings {
	return GameSettings{
		BoardWidth:  4,
		BoardHeight: 4,
	}
}
