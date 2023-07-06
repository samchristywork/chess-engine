package main

import (
  "fmt"
  "net/http"
)

var globalBoard Board

type Piece byte

type Square struct {
  File int
  Rank int
}

type Board struct {
  board [8][8]Piece
  activeColor string
  castling string
  enPassant string
  halfMoveClock int
  fullMoveNumber int
  lastMove [2]string
  moveList []string
  whiteAI func(*Board)
  blackAI func(*Board)
}

func main() {
  resetHandler(nil, nil)

  http.HandleFunc("/board", boardHandler)
  http.HandleFunc("/move", movePieceHandler)
  http.HandleFunc("/valid-moves", validMoveHandler)
  http.HandleFunc("/fen", currentFENHandler)
  http.HandleFunc("/reset", resetHandler)
  http.HandleFunc("/move-list", moveListHandler)
  http.HandleFunc("/computer-move", computerMoveHandler)
  http.HandleFunc("/piece-values", pieceValuesHandler)

  fmt.Printf("Starting server at port 8080\n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Printf("Error: %s\n", err)
  }
}
