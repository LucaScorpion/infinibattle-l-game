package parser

type turnState struct {
	GameState gameState `json:"GameState"`
	Turn      int       `json:"Turn"`
	Player    int       `json:"Player"` // 1 or 0
}

type gameState struct {
	Board        boardState `json:"Board"`
	ScorePlayer0 int        `json:"ScorePlayer0"`
	ScorePlayer1 int        `json:"ScorePlayer1"`
}

type boardState struct {
	Board [4][4]pieceType `json:"Board"`
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
	PlayerLPieceCoordinates [4]coordinate `json:"PlayerLPieceCoordinates"`
	NeutralPieceCoordinates [2]coordinate `json:"NeutralPieceCoordinates"`
}
