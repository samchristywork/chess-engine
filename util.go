package main;

import (
  "fmt"
  "strings"
)

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func printBoard(board Board) {
  fmt.Println("    a b c d e f g h")
  fmt.Println("   ----------------")
  for n, rank := range board.board {
    fmt.Printf("%d |", 8-n)
    for _, piece := range rank {
      fmt.Printf(" %c", piece)
    }
    fmt.Println()
  }

  fen := boardToFEN(board)
  fmt.Println(fen)
}

func printValidMoves(board Board, square string) {
  fmt.Println("    a b c d e f g h")
  fmt.Println("   ----------------")
  fromSquare := parseSquare(square)
  for rank := range board.board {
    fmt.Printf("%d |", 8-rank)
    for file := range board.board[rank] {
      toSquare := Square{file, rank}

      if getPiece(board, toSquare) == Piece(' ') {
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

func listAllMoves(board Board) [][2]string {
  moves := [][2]string{}
  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      fromSquare := Square{file, rank}
      piece := board.board[rank][file]
      if (isWhite(piece) && board.activeColor == "w") || (isBlack(piece) && board.activeColor == "b") {
        for rank := 0; rank < 8; rank++ {
          for file := 0; file < 8; file++ {
            toSquare := Square{file, rank}
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

func pieceHasMoves(board Board, fromSquare Square) Square {
  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      toSquare := Square{file, rank}
      if isValidMove(&board, fromSquare, toSquare) {
        return toSquare
      }
    }
  }

  return Square{-1, -1}
}

func listAllPiecesWithMoves(board Board) []string {
  pieces := []string{}
  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      fromSquare := Square{file, rank}
      piece := board.board[rank][file]
      if (isWhite(piece) && board.activeColor == "w") || (isBlack(piece) && board.activeColor == "b") {
        oneMove := pieceHasMoves(board, fromSquare);
        if (oneMove.File != -1) {
          pieces = append(pieces, fmt.Sprintf("%c%c", 'a' + fromSquare.File, '8' - fromSquare.Rank))
        }
      }
    }
  }
  return pieces
}

func parseSquare(square string) Square {
  file := int(square[0] - 'a')
  rank := 8 - int(square[1] - '0')
  return Square{file, rank}
}

func isOutOfBounds(square Square) bool {
  return square.File < 0 || square.File > 7 || square.Rank < 0 || square.Rank > 7
}

func isSameColor(piece1 Piece, piece2 Piece) bool {
  if (isWhite(piece1) && isWhite(piece2)) {
    return true
  }
  if (isBlack(piece1) && isBlack(piece2)) {
    return true
  }
  return false
}

func movePiece(board *Board, from string, to string) bool {
  fromSquare := parseSquare(from)
  toSquare := parseSquare(to)

  if isOutOfBounds(fromSquare) || isOutOfBounds(toSquare) {
    return false
  }

  if getPiece(*board, fromSquare) == Piece(' ') {
    return false
  }

  if !isValidMove(board, fromSquare, toSquare) {
    return false
  }

  board.board[toSquare.Rank][toSquare.File] = board.board[fromSquare.Rank][fromSquare.File]
  board.board[fromSquare.Rank][fromSquare.File] = Piece(' ')

  board.moveList = append(board.moveList, fmt.Sprintf("%s-%s", from, to))

  if (board.activeColor == "w") {
    board.activeColor = "b"
  } else {
    board.activeColor = "w"
  }

  board.lastMove[0] = from
  board.lastMove[1] = to

  return true
}

func isWhite(piece Piece) bool {
  return piece >= 'A' && piece <= 'Z'
}

func isBlack(piece Piece) bool {
  return piece >= 'a' && piece <= 'z'
}

func getRankFile(board Board, rank int, file int) Piece {
  return getPiece(board, Square{file, rank})
}

func getPiece(board Board, square Square) Piece {
  if isOutOfBounds(square) {
    return Piece(' ')
  }
  return board.board[square.Rank][square.File]
}

func setPiece(board *Board, square string, piece Piece) {
  square = strings.ToLower(square)
  sq := parseSquare(square)
  board.board[sq.Rank][sq.File] = piece
}

func pieceToEmoji(piece Piece) string {
  switch piece {
  case Piece(' '):
    return " "
  case Piece('P'):
    return "♙"
  case Piece('N'):
    return "♘"
  case Piece('B'):
    return "♗"
  case Piece('R'):
    return "♖"
  case Piece('Q'):
    return "♕"
  case Piece('K'):
    return "♔"
  case Piece('p'):
    return "♟"
  case Piece('n'):
    return "♞"
  case Piece('b'):
    return "♝"
  case Piece('r'):
    return "♜"
  case Piece('q'):
    return "♛"
  case Piece('k'):
    return "♚"
  }
  return " "
}
