package main

import (
	"testing"
)

func Test_makeHTMLBoard(t *testing.T) {
	board := FENToBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")

	html := makeHTMLBoard(false, board);

	expect := 5561
	if len(html) != expect {
		t.Fatalf("got %d, expected %d.", len(html), expect)
	}
}
