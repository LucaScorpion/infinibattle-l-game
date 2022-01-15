package parser

type turnState struct {
	GameState gameState `json:"gameState"`
	Turn      int       `json:"turn"`
	Player    int       `json:"player"` // One of the piece types.
}

type gameState struct {
	Board        boardState `json:"board"`
	ScorePlayer0 int        `json:"scorePlayer0"`
	ScorePlayer1 int        `json:"scorePlayer1"`
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
