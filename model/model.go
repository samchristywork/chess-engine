package model

type Piece byte

type Square struct {
	File int
	Rank int
}

type Board struct {
	Board          [8][8]Piece
	ActiveColor    string
	Castling       string
	EnPassant      string
	HalfMoveClock  int
	FullMoveNumber int
	LastMove       [2]string
	MoveList       []string
	WhiteAI        func(*Board)
	BlackAI        func(*Board)
}
