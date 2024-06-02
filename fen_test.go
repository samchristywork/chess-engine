package main

import (
	"testing"
)

func Test_FENToBoard(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fen := boardToFEN(board)

	expect := "rnbqkbnr/pppppppp/////PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	if fen != expect {
		t.Fatalf("got %s, expected %s.", fen, expect)
	}
}
func Test_boardToFEN(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	fen := boardToFEN(board)

	expect := "rnbqkbnr/pppppppp/////PPPPPPPP/RNBQKBNR w KQkq - 0 1"
	if fen != expect {
		t.Fatalf("got %s, expected %s.", fen, expect)
	}
}
