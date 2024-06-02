package main

import (
	"chess-engine/model"
	"testing"
)

func Test_abs(t *testing.T) {
	v := abs(-1)

	expect := 1
	if v != expect {
		t.Fatalf("got %d, expected %d.", v, expect)
	}
}

func Test_printBoard(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	printBoard(board)
}

func Test_printValidMoves(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	printValidMoves(board, "a1")
}

func Test_listAllMoves(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	moves := listAllMoves(board)

	expect := 20
	if len(moves) != expect {
		t.Fatalf("got %d, expected %d.", len(moves), expect)
	}
}

func Test_pieceHasMoves(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fromSquare := model.Square{File: 0, Rank: 1}
	hasMoves := pieceHasMoves(board, fromSquare)

	expect := true
	if hasMoves != expect {
		t.Fatalf("got %t, expected %t.", hasMoves, expect)
	}
}

func Test_listAllPiecesWithMoves(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	pieces := listAllPiecesWithMoves(board)

	expect := 10
	if len(pieces) != expect {
		t.Fatalf("got %d, expected %d.", len(pieces), expect)
	}
}

func Test_parseSquare(t *testing.T) {
	square := parseSquare("a2")

	expect := 0
	if square.File != expect {
		t.Fatalf("got %d, expected %d.", square.File, expect)
	}

	expect = 6
	if square.Rank != expect {
		t.Fatalf("got %d, expected %d.", square.Rank, expect)
	}
}

func Test_isOutOfBounds(t *testing.T) {
	square := model.Square{File: -1, Rank: 1}
	oob := isOutOfBounds(square)

	expect := true
	if oob != expect {
		t.Fatalf("got %t, expected %t.", oob, expect)
	}
}

func Test_isSameColor(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	a := board.Board[0][0]
	b := board.Board[7][0]
	sameColor := isSameColor(a, b)

	expect := false
	if sameColor != expect {
		t.Fatalf("got %t, expected %t.", sameColor, expect)
	}

	a = board.Board[0][0]
	b = board.Board[0][7]
	sameColor = isSameColor(a, b)

	expect = true
	if sameColor != expect {
		t.Fatalf("got %t, expected %t.", sameColor, expect)
	}
}

func Test_movePiece(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	moved := movePiece(&board, "b7", "b5")
	printBoard(board)
	fen := boardToFEN(board)

	expect := "rnbqkbnr/pppppppp///P//PPPPPPP/RNBQKBNR b KQkq - 0 1"
	if fen != expect {
		t.Fatalf("got %s, expected %s.", fen, expect)
	}

	expect2 := true
	if moved != expect2 {
		t.Fatalf("got %t, expected %t.", moved, expect2)
	}
}

func Test_isWhite(t *testing.T) {
	piece := model.Piece('B')
	iw := isWhite(piece)

	expect := true
	if iw != expect {
		t.Fatalf("got %t, expected %t.", iw, expect)
	}
}

func Test_isBlack(t *testing.T) {
	piece := model.Piece('B')
	ib := isBlack(piece)

	expect := false
	if ib != expect {
		t.Fatalf("got %t, expected %t.", ib, expect)
	}
}

func Test_getRankFile(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	piece := getRankFile(board, 0, 1)

	expect := model.Piece('N')
	if piece != expect {
		t.Fatalf("got %s, expected %s.", string(piece), string(expect))
	}
}

func Test_getPiece(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	square := model.Square{File: 1, Rank: 0}
	piece := getPiece(board, square)

	expect := model.Piece('N')
	if piece != expect {
		t.Fatalf("got %s, expected %s.", string(piece), string(expect))
	}
}

func Test_setPiece(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	setPiece(&board, "a6", model.Piece('Q'))

	square := parseSquare("a6")
	piece := getPiece(board, square)

	expect := model.Piece('Q')
	if piece != expect {
		t.Fatalf("got %s, expected %s.", string(piece), string(expect))
	}
}

func Test_pieceToEmoji(t *testing.T) {
	piece := pieceToEmoji('P')

	expect := "♙"
	if piece != expect {
		t.Fatalf("got %s, expected %s.", string(piece), string(expect))
	}
}
