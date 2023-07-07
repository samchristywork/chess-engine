package main

func isValidPawnMove(board *Board, from Square, to Square) bool {
  piece := board.board[from.Rank][from.File]
  target := board.board[to.Rank][to.File]

  fileDiff := abs(from.File - to.File)
  rowDiff := abs(from.Rank - to.Rank)
  isStartingRank := (isWhite(piece) && from.Rank == 1) || (isBlack(piece) && from.Rank == 6)

  if fileDiff > 1 {
    return false
  }

  if !isStartingRank && rowDiff > 1 {
    return false
  }

  if fileDiff == 1 && rowDiff == 1 {
    if target == Piece(' ') {
      return false
    }
  }

  if fileDiff == 0 && rowDiff == 1 {
    if target != Piece(' ') {
      return false
    }
  }

  if rowDiff == 2 && fileDiff != 0 {
    return false
  }

  if rowDiff == 2 && board.board[(from.Rank + to.Rank) / 2][from.File] != Piece(' ') {
    return false
  }

  if fileDiff == 0 && target != Piece(' ') {
    return false
  }

  if rowDiff > 2 {
    return false
  }

  // Invalid if color of the piece and the target are the same
  if (isSameColor(piece, target)) {
    return false;
  }

  // Piece can only move forward
  if isWhite(piece) {
    if (to.Rank <= from.Rank) {
      return false;
    }
  } else {
    if (to.Rank >= from.Rank) {
      return false;
    }
  }

  return true;
}
