package main;

import (
  "chess-engine/model"
  "fmt"
  "math/rand"
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

func evaluateWhite(board model.Board) int {
  score := 0

  for rank := range board.Board {
    for file := range board.Board[rank] {
      square := model.Square{file, rank}
      piece := getPiece(board, square)

      if piece == model.Piece(' ') {
        continue
      }

      if piece == model.Piece('P') {
        score += 100
        score += pawnTable[rank][file]
      } else if piece == model.Piece('N') {
        score += 320
        score += knightTable[rank][file]
      } else if piece == model.Piece('B') {
        score += 330
        score += bishopTable[rank][file]
      } else if piece == model.Piece('R') {
        score += 500
        score += rookTable[rank][file]
      } else if piece == model.Piece('Q') {
        score += 900
        score += queenTable[rank][file]
      } else if piece == model.Piece('K') {
        score += 20000
        score += kingTable[rank][file]
      }
    }
  }

  return score
}

func evaluateBlack(board model.Board) int {
  score := 0

  for rank := range board.Board {
    for file := range board.Board[rank] {
      square := model.Square{file, rank}
      piece := getPiece(board, square)

      if piece == model.Piece(' ') {
        continue
      }

      if piece == model.Piece('p') {
        score += 100
        score += pawnTable[rank][file]
      } else if piece == model.Piece('n') {
        score += 320
        score += knightTable[rank][file]
      } else if piece == model.Piece('b') {
        score += 330
        score += bishopTable[rank][file]
      } else if piece == model.Piece('r') {
        score += 500
        score += rookTable[rank][file]
      } else if piece == model.Piece('q') {
        score += 900
        score += queenTable[rank][file]
      } else if piece == model.Piece('k') {
        score += 20000
        score += kingTable[rank][file]
      }
    }
  }

  return score
}

func evaluateBoard(board model.Board) int {
  if board.ActiveColor == "w" {
    return evaluateWhite(board) - evaluateBlack(board)
  } else {
    return evaluateBlack(board) - evaluateWhite(board)
  }
}

func minimax(board model.Board, depth int, maximizing bool) int {
  if depth == 0 {
    return evaluateBoard(board)
  }

  moves := listAllMoves(board)

  if len(moves) == 0 {
    return evaluateBoard(board)
  }

  if maximizing {
    bestValue := -9999
    for _, move := range moves {
      newBoard := board
      movePiece(&newBoard, move[0], move[1])
      value := minimax(newBoard, depth - 1, false)
      if value > bestValue {
        bestValue = value
      }
    }
    return bestValue
  } else {
    bestValue := 9999
    for _, move := range moves {
      newBoard := board
      movePiece(&newBoard, move[0], move[1])
      value := minimax(newBoard, depth - 1, true)
      if value < bestValue {
        bestValue = value
      }
    }
    return bestValue
  }
}

func minimaxABPruning(board model.Board, depth int, alpha int, beta int, maximizing bool) int {
  if depth == 0 {
    return evaluateBoard(board)
  }

  moves := listAllMoves(board)

  if len(moves) == 0 {
    return evaluateBoard(board)
  }

  if maximizing {
    bestValue := -9999
    for _, move := range moves {
      newBoard := board
      movePiece(&newBoard, move[0], move[1])
      value := minimaxABPruning(newBoard, depth - 1, alpha, beta, false)
      if value > bestValue {
        bestValue = value
      }
      if value > alpha {
        alpha = value
      }
      if alpha >= beta {
        break
      }
    }
    return bestValue
  } else {
    bestValue := 9999
    for _, move := range moves {
      newBoard := board
      movePiece(&newBoard, move[0], move[1])
      value := minimaxABPruning(newBoard, depth - 1, alpha, beta, true)
      if value < bestValue {
        bestValue = value
      }
      if value < beta {
        beta = value
      }
      if alpha >= beta {
        break
      }
    }
    return bestValue
  }
}

func isMoveValid(board model.Board, from string, to string) int {
  fromSquare := parseSquare(from)
  toSquare := parseSquare(to)

  if isOutOfBounds(fromSquare) || isOutOfBounds(toSquare) {
    return 0
  }

  if getPiece(board, fromSquare) == model.Piece(' ') {
    return 0
  }

  if !isValidMove(&board, fromSquare, toSquare) {
    return 0
  }

  return 1
}

type Move struct {
  from string
  to string
  score int
}

func computeMinimaxABPruningMove(board *model.Board) {
  moves := listAllMoves(*board)
  validMoves := make([]Move, 0)

  for _, move := range moves {
    if isMoveValid(*board, move[0], move[1]) == 1 {
      newBoard := *board
      movePiece(&newBoard, move[0], move[1])
      score := minimaxABPruning(newBoard, 4, -9999, 9999, true)
      fmt.Printf("%s -> %s: %d\n", move[0], move[1], score)
      validMoves = append(validMoves, Move{move[0], move[1], score})
    }
  }

  if len(validMoves) == 0 {
    return
  }

  fmt.Println(validMoves)

  move := validMoves[0]
  for _, validMove := range validMoves {
    if validMove.score < move.score {
      move = validMove
    }
  }

  movePiece(board, move.from, move.to)
}

func computeMinimaxMove(board *model.Board) {
  moves := listAllMoves(*board)
  validMoves := make([]Move, 0)

  for _, move := range moves {
    if isMoveValid(*board, move[0], move[1]) == 1 {
      newBoard := *board
      movePiece(&newBoard, move[0], move[1])
      score := minimax(newBoard, 3, true)
      validMoves = append(validMoves, Move{move[0], move[1], score})
    }
  }

  if len(validMoves) == 0 {
    return
  }

  fmt.Println(validMoves)

  move := validMoves[0]
  for _, validMove := range validMoves {
    if validMove.score > move.score {
      move = validMove
    }
  }

  movePiece(board, move.from, move.to)
}

func computeRandomMove(board *model.Board) {
  moves := listAllMoves(*board)

  if len(moves) == 0 {
    return
  }

  move := moves[rand.Intn(len(moves))]

  movePiece(board, move[0], move[1])
}
