package main

import (
  "fmt"
  "strings"
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
}
