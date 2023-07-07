package main

import (
  "fmt"
  "strconv"
  "strings"
)

func FENToBoard(fen string) Board {
  board := Board{}
  for rank := range board.board {
    for file := range board.board[rank] {
      board.board[rank][file] = Piece(' ')
    }
  }

  fenParts := strings.Split(fen, " ")

  boardPart := fenParts[0]
  activeColor := fenParts[1]
  castling := fenParts[2]
  enPassant := fenParts[3]
  halfMoveClock := fenParts[4]
  fullMoveNumber := fenParts[5]

  board.activeColor = activeColor
  board.castling = castling
  board.enPassant = enPassant
  board.halfMoveClock, _ = strconv.Atoi(halfMoveClock)
  board.fullMoveNumber, _ = strconv.Atoi(fullMoveNumber)

  ranks := strings.Split(boardPart, "/")
  for rank, fenRank := range ranks {
    file := 0
    for _, rune := range fenRank {
      if rune >= '1' && rune <= '8' {
        file += int(rune - '0')
      } else {
        board.board[7-rank][file] = Piece(rune)
        file++
      }
    }
  }

  return board
}

func boardToFEN(board Board) string {
  fen := ""
  for rank := 7; rank >= 0; rank-- {
    emptySquares := 0
    for file := 0; file < 8; file++ {
      piece := board.board[rank][file]
      if piece == Piece(' ') {
        emptySquares++
      } else {
        if emptySquares > 0 {
          fen += fmt.Sprintf("%d", emptySquares)
          emptySquares = 0
        }
        fen += string(piece)
      }
    }
    if emptySquares > 0 {
      fen += fmt.Sprintf("%d", emptySquares)
    }
    if rank > 0 {
      fen += "/"
    }
  }

  fen += fmt.Sprintf(" %s %s %s %d %d", board.activeColor, board.castling, board.enPassant, board.halfMoveClock, board.fullMoveNumber)

  return fen
}
