package main

import (
	"testing"
)

func Test_isValidPawnMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	from := parseSquare("a7")
	to := parseSquare("a5")
	valid := isValidPawnMove(&board, from, to)

	expect := true
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}

func Test_isValidRookMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	from := parseSquare("a8")
	to := parseSquare("a5")
	valid := isValidRookMove(&board, from, to)

	expect := false
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}

func Test_isValidKnightMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	from := parseSquare("b8")
	to := parseSquare("a6")
	valid := isValidKnightMove(&board, from, to)

	expect := true
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}

func Test_isValidBishopMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	from := parseSquare("c8")
	to := parseSquare("a6")
	valid := isValidBishopMove(&board, from, to)

	expect := false
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}

func Test_isValidQueenMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	from := parseSquare("d8")
	to := parseSquare("a6")
	valid := isValidQueenMove(&board, from, to)

	expect := false
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}

func Test_isValidKingMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	from := parseSquare("e8")
	to := parseSquare("a6")
	valid := isValidKingMove(&board, from, to)

	expect := false
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}

func Test_isValidMove(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	from := parseSquare("g8")
	to := parseSquare("h6")
	valid := isValidMove(&board, from, to)

	expect := true
	if valid != expect {
		t.Fatalf("got %t, expected %t.", valid, expect)
	}
}
