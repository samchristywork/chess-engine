package main

import (
	"testing"
)

func Test_evaluateWhite(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	score := evaluateWhite(board)

	expect := 24245
	if score != expect {
		t.Fatalf("got %d, expected %d.", score, expect)
	}
}

func Test_evaluateBlack(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	score := evaluateBlack(board)

	expect := 23905
	if score != expect {
		t.Fatalf("got %d, expected %d.", score, expect)
	}
}

func Test_evaluateBoard(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	score := evaluateBoard(board)

	expect := 340
	if score != expect {
		t.Fatalf("got %d, expected %d.", score, expect)
	}

	board.ActiveColor = "b"

	score = evaluateBoard(board)

	expect = -expect
	if score != expect {
		t.Fatalf("got %d, expected %d.", score, expect)
	}
}

func Test_minimax(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	best := minimax(board, 3, true)

	expect := -270
	if best != expect {
		t.Fatalf("got %d, expected %d.", best, expect)
	}
}

func Test_minimaxABPruning(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	best := minimaxABPruning(board, 3, -9999, 9999, true)

	expect := -270
	if best != expect {
		t.Fatalf("got %d, expected %d.", best, expect)
	}
}

func Test_isMoveValid(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	moves := listAllMoves(board)

	valid := isMoveValid(board, moves[0][0], moves[0][1])

	expect := true
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}

func Test_computeMinimaxABPruningMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	computeMinimaxABPruningMove(&board)

	fen := boardToFEN(board)

	expect := "rnbqkbnr/pppppppp////N/PPPPPPPP/RBQKBNR b KQkq - 0 1"
	if fen != expect {
		t.Fatalf("got %s, expected %s.", fen, expect)
	}
}

func Test_computeMinimaxMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	computeMinimaxMove(&board)

	fen := boardToFEN(board)

	expect := "rnbqkbnr/pppppppp////N/PPPPPPPP/RBQKBNR b KQkq - 0 1"
	if fen != expect {
		t.Fatalf("got %s, expected %s.", fen, expect)
	}
}
