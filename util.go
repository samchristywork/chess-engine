package main;

import (
  "chess-engine/model"
  "fmt"
  "strings"
)

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func printBoard(board model.Board) {
  fmt.Println("    a b c d e f g h")
  fmt.Println("   ----------------")
  for n, rank := range board.Board {
    fmt.Printf("%d |", 8-n)
    for _, piece := range rank {
      fmt.Printf(" %c", piece)
    }
    fmt.Println()
  }

  fen := boardToFEN(board)
  fmt.Println(fen)
}

func printValidMoves(board model.Board, square string) {
  fmt.Println("    a b c d e f g h")
  fmt.Println("   ----------------")
  fromSquare := parseSquare(square)
  for rank := range board.Board {
    fmt.Printf("%d |", 8-rank)
    for file := range board.Board[rank] {
      toSquare := model.Square{file, rank}

      if getPiece(board, toSquare) == model.Piece(' ') {
        if isValidMove(&board, fromSquare, toSquare) {
          fmt.Print(" .")
        } else {
          fmt.Print("  ")
        }
      } else {
        if isValidMove(&board, fromSquare, toSquare) {
          fmt.Print(" c")
        } else {
          fmt.Printf(" %c", getPiece(board, toSquare))
        }
      }
    }
    fmt.Println()
  }
}

func listAllMoves(board model.Board) [][2]string {
  moves := [][2]string{}
  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      fromSquare := model.Square{file, rank}
      piece := board.Board[rank][file]
      if (isWhite(piece) && board.ActiveColor == "w") || (isBlack(piece) && board.ActiveColor == "b") {
        for rank := 0; rank < 8; rank++ {
          for file := 0; file < 8; file++ {
            toSquare := model.Square{file, rank}
            if isValidMove(&board, fromSquare, toSquare) {
              moves = append(moves, [2]string{fmt.Sprintf("%c%c", 'a' + fromSquare.File, '8' - fromSquare.Rank), fmt.Sprintf("%c%c", 'a' + toSquare.File, '8' - toSquare.Rank)})
            }
          }
        }
      }
    }
  }
  return moves
}

func pieceHasMoves(board model.Board, fromSquare model.Square) model.Square {
  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      toSquare := model.Square{file, rank}
      if isValidMove(&board, fromSquare, toSquare) {
        return toSquare
      }
    }
  }

  return model.Square{-1, -1}
}

func listAllPiecesWithMoves(board model.Board) []string {
  pieces := []string{}
  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      fromSquare := model.Square{file, rank}
      piece := board.Board[rank][file]
      if (isWhite(piece) && board.ActiveColor == "w") || (isBlack(piece) && board.ActiveColor == "b") {
        oneMove := pieceHasMoves(board, fromSquare);
        if (oneMove.File != -1) {
          pieces = append(pieces, fmt.Sprintf("%c%c", 'a' + fromSquare.File, '8' - fromSquare.Rank))
        }
      }
    }
  }
  return pieces
}

func parseSquare(square string) model.Square {
  file := int(square[0] - 'a')
  rank := 8 - int(square[1] - '0')
  return model.Square{file, rank}
}

func isOutOfBounds(square model.Square) bool {
  return square.File < 0 || square.File > 7 || square.Rank < 0 || square.Rank > 7
}

func isSameColor(piece1 model.Piece, piece2 model.Piece) bool {
  if (isWhite(piece1) && isWhite(piece2)) {
    return true
  }
  if (isBlack(piece1) && isBlack(piece2)) {
    return true
  }
  return false
}

func movePiece(board *model.Board, from string, to string) bool {
  fromSquare := parseSquare(from)
  toSquare := parseSquare(to)

  if isOutOfBounds(fromSquare) || isOutOfBounds(toSquare) {
    return false
  }

  if getPiece(*board, fromSquare) == model.Piece(' ') {
    return false
  }

  if !isValidMove(board, fromSquare, toSquare) {
    return false
  }

  board.Board[toSquare.Rank][toSquare.File] = board.Board[fromSquare.Rank][fromSquare.File]
  board.Board[fromSquare.Rank][fromSquare.File] = model.Piece(' ')

  board.MoveList = append(board.MoveList, fmt.Sprintf("%s-%s", from, to))

  if (board.ActiveColor == "w") {
    board.ActiveColor = "b"
  } else {
    board.ActiveColor = "w"
  }

  board.LastMove[0] = from
  board.LastMove[1] = to

  return true
}

func isWhite(piece model.Piece) bool {
  return piece >= 'A' && piece <= 'Z'
}

func isBlack(piece model.Piece) bool {
  return piece >= 'a' && piece <= 'z'
}

func getRankFile(board model.Board, rank int, file int) model.Piece {
  return getPiece(board, model.Square{file, rank})
}

func getPiece(board model.Board, square model.Square) model.Piece {
  if isOutOfBounds(square) {
    return model.Piece(' ')
  }
  return board.Board[square.Rank][square.File]
}

func setPiece(board *model.Board, square string, piece model.Piece) {
  square = strings.ToLower(square)
  sq := parseSquare(square)
  board.Board[sq.Rank][sq.File] = piece
}

func pieceToEmoji(piece model.Piece) string {
  switch piece {
  case model.Piece(' '):
    return " "
  case model.Piece('P'):
    return "♙"
  case model.Piece('N'):
    return "♘"
  case model.Piece('B'):
    return "♗"
  case model.Piece('R'):
    return "♖"
  case model.Piece('Q'):
    return "♕"
  case model.Piece('K'):
    return "♔"
  case model.Piece('p'):
    return "♟"
  case model.Piece('n'):
    return "♞"
  case model.Piece('b'):
    return "♝"
  case model.Piece('r'):
    return "♜"
  case model.Piece('q'):
    return "♛"
  case model.Piece('k'):
    return "♚"
  }
  return " "
}
