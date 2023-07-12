package main;

import (
  "math/rand"
  "fmt"
)

var pawnTable = [8][8]int{
  {0, 0, 0, 0, 0, 0, 0, 0},
  {50, 50, 50, 50, 50, 50, 50, 50},
  {10, 10, 20, 30, 30, 20, 10, 10},
  {5, 5, 10, 25, 25, 10, 5, 5},
  {0, 0, 0, 20, 20, 0, 0, 0},
  {5, -5, -10, 0, 0, -10, -5, 5},
  {5, 10, 10, -20, -20, 10, 10, 5},
  {0, 0, 0, 0, 0, 0, 0, 0},
}

var knightTable = [8][8]int{
  {-50, -40, -30, -30, -30, -30, -40, -50},
  {-40, -20, 0, 0, 0, 0, -20, -40},
  {-30, 0, 10, 15, 15, 10, 0, -30},
  {-30, 5, 15, 20, 20, 15, 5, -30},
  {-30, 0, 15, 20, 20, 15, 0, -30},
  {-30, 5, 10, 15, 15, 10, 5, -30},
  {-40, -20, 0, 5, 5, 0, -20, -40},
  {-50, -40, -30, -30, -30, -30, -40, -50},
}

var bishopTable = [8][8]int{
  {-20, -10, -10, -10, -10, -10, -10, -20},
  {-10, 0, 0, 0, 0, 0, 0, -10},
  {-10, 0, 5, 10, 10, 5, 0, -10},
  {-10, 5, 5, 10, 10, 5, 5, -10},
  {-10, 0, 10, 10, 10, 10, 0, -10},
  {-10, 10, 10, 10, 10, 10, 10, -10},
  {-10, 5, 0, 0, 0, 0, 5, -10},
  {-20, -10, -10, -10, -10, -10, -10, -20},
}

var rookTable = [8][8]int{
  {0, 0, 0, 0, 0, 0, 0, 0},
  {5, 10, 10, 10, 10, 10, 10, 5},
  {-5, 0, 0, 0, 0, 0, 0, -5},
  {-5, 0, 0, 0, 0, 0, 0, -5},
  {-5, 0, 0, 0, 0, 0, 0, -5},
  {-5, 0, 0, 0, 0, 0, 0, -5},
  {-5, 0, 0, 0, 0, 0, 0, -5},
  {0, 0, 0, 5, 5, 0, 0, 0},
}

var queenTable = [8][8]int{
  {-20, -10, -10, -5, -5, -10, -10, -20},
  {-10, 0, 0, 0, 0, 0, 0, -10},
  {-10, 0, 5, 5, 5, 5, 0, -10},
  {-5, 0, 5, 5, 5, 5, 0, -5},
  {0, 0, 5, 5, 5, 5, 0, -5},
  {-10, 5, 5, 5, 5, 5, 0, -10},
  {-10, 0, 5, 0, 0, 0, 0, -10},
  {-20, -10, -10, -5, -5, -10, -10, -20},
}

var kingTable = [8][8]int{
  {-30, -40, -40, -50, -50, -40, -40, -30},
  {-30, -40, -40, -50, -50, -40, -40, -30},
  {-30, -40, -40, -50, -50, -40, -40, -30},
  {-30, -40, -40, -50, -50, -40, -40, -30},
  {-20, -30, -30, -40, -40, -30, -30, -20},
  {-10, -20, -20, -20, -20, -20, -20, -10},
  {20, 20, 0, 0, 0, 0, 20, 20},
  {20, 30, 10, 0, 0, 10, 30, 20},
}

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
        score += pawnTable[rank][file]
      } else if piece == Piece('N') {
        score += 320
        score += knightTable[rank][file]
      } else if piece == Piece('B') {
        score += 330
        score += bishopTable[rank][file]
      } else if piece == Piece('R') {
        score += 500
        score += rookTable[rank][file]
      } else if piece == Piece('Q') {
        score += 900
        score += queenTable[rank][file]
      } else if piece == Piece('K') {
        score += 20000
        score += kingTable[rank][file]
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
        score += pawnTable[rank][file]
      } else if piece == Piece('n') {
        score += 320
        score += knightTable[rank][file]
      } else if piece == Piece('b') {
        score += 330
        score += bishopTable[rank][file]
      } else if piece == Piece('r') {
        score += 500
        score += rookTable[rank][file]
      } else if piece == Piece('q') {
        score += 900
        score += queenTable[rank][file]
      } else if piece == Piece('k') {
        score += 20000
        score += kingTable[rank][file]
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
