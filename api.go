package main

import (
  "fmt"
  "net/http"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
  html := ""
  html += makeHTMLBoard(false, globalBoard)
  html += makeHTMLBoard(true, globalBoard)

  fmt.Fprintf(w, "%s\n", html)
}

func validMoveHandler(w http.ResponseWriter, r *http.Request) {
  from := r.URL.Query().Get("from")

  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      toSquare := Square{file, rank}
      if isValidMove(&globalBoard, parseSquare(from), toSquare) {
        fmt.Fprintf(w, "%c%c\n", 'a' + file, '8' - rank)
      }
    }
  }
}

func currentFENHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "%s\n", boardToFEN(globalBoard))
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  globalBoard = FENToBoard(fen)
  globalBoard.moveList = []string{}
}

func moveListHandler(w http.ResponseWriter, r *http.Request) {
  for n, move := range globalBoard.moveList {
    fmt.Fprintf(w, "<div class='move-list-item'>%d: %s</div>", n + 1, move)
  }
}

func computerMoveHandler(w http.ResponseWriter, r *http.Request) {
}

func movePieceHandler(w http.ResponseWriter, r *http.Request) {
}

var pieceValueMap = map[rune]int{
  'P': 1,
  'p': 1,
  'R': 5,
  'r': 5,
  'N': 3,
  'n': 3,
  'B': 3,
  'b': 3,
  'Q': 9,
  'q': 9,
  'K': 0,
  'k': 0,
}

func pieceValuesHandler(w http.ResponseWriter, r *http.Request) {
}

