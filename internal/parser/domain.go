package parser

type turnState struct {
	GameState gameState `json:"gameState"`
	Turn      int       `json:"turn"`
	Player    int       `json:"player"`
}

type gameState struct {
	Board boardState `json:"board"`
}

type boardState struct {
	Board [4][4]pieceType `json:"board"`
}

type pieceType int

const (
	empty   pieceType = 0
	player1 pieceType = 1
	player2 pieceType = 2
	neutral pieceType = 4
)

type coordinate [2]int

type placePiecesCommand struct {
	PlayerLPieceCoordinates [4]coordinate `json:"playerLPieceCoordinates"`
	NeutralPieceCoordinates [2]coordinate `json:"neutralPieceCoordinates"`
}
