package main;

import (
  "math/rand"
  "fmt"
)

func evaluateWhite(board Board) int {
  score := 0

  for rank := range board.board {
    for file := range board.board[rank] {
      square := Square{file, rank}
      piece := getPiece(board, square)

      if piece == Piece(' ') {
        continue
      }

      if piece == Piece('P') {
        score += 100
      } else if piece == Piece('N') {
        score += 320
      } else if piece == Piece('B') {
        score += 330
      } else if piece == Piece('R') {
        score += 500
      } else if piece == Piece('Q') {
        score += 900
      } else if piece == Piece('K') {
        score += 20000
      }
    }
  }

  return score
}

func evaluateBlack(board Board) int {
  score := 0

  for rank := range board.board {
    for file := range board.board[rank] {
      square := Square{file, rank}
      piece := getPiece(board, square)

      if piece == Piece(' ') {
        continue
      }

      if piece == Piece('p') {
        score += 100
      } else if piece == Piece('n') {
        score += 320
      } else if piece == Piece('b') {
        score += 330
      } else if piece == Piece('r') {
        score += 500
      } else if piece == Piece('q') {
        score += 900
      } else if piece == Piece('k') {
        score += 20000
      }
    }
  }

  return score
}

func evaluateBoard(board Board) int {
  if board.activeColor == "w" {
    return evaluateWhite(board) - evaluateBlack(board)
  } else {
    return evaluateBlack(board) - evaluateWhite(board)
  }
}

func computeRandomMove(board *Board) {
  moves := listAllMoves(*board)

  if len(moves) == 0 {
    return
  }

  move := moves[rand.Intn(len(moves))]

  movePiece(board, move[0], move[1])
}
