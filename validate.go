package main

import (
	"chess-engine/model"
)

type ValidationFunc func(*model.Board, model.Square, model.Square) bool

var pieceValidationMap = map[rune]ValidationFunc{
	'P': isValidPawnMove,
	'p': isValidPawnMove,
	'R': isValidRookMove,
	'r': isValidRookMove,
	'N': isValidKnightMove,
	'n': isValidKnightMove,
	'B': isValidBishopMove,
	'b': isValidBishopMove,
	'Q': isValidQueenMove,
	'q': isValidQueenMove,
	'K': isValidKingMove,
	'k': isValidKingMove,
}

func isValidPawnMove(board *model.Board, from model.Square, to model.Square) bool {
	piece := board.Board[from.Rank][from.File]
	target := board.Board[to.Rank][to.File]

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
		if target == model.Piece(' ') {
			return false
		}
	}

	if fileDiff == 0 && rowDiff == 1 {
		if target != model.Piece(' ') {
			return false
		}
	}

	if rowDiff == 2 && fileDiff != 0 {
		return false
	}

	if rowDiff == 2 && board.Board[(from.Rank+to.Rank)/2][from.File] != model.Piece(' ') {
		return false
	}

	if fileDiff == 0 && target != model.Piece(' ') {
		return false
	}

	if rowDiff > 2 {
		return false
	}

	// Invalid if color of the piece and the target are the same
	if isSameColor(piece, target) {
		return false
	}

	// Piece can only move forward
	if isWhite(piece) {
		if to.Rank <= from.Rank {
			return false
		}
	} else {
		if to.Rank >= from.Rank {
			return false
		}
	}

	return true
}

func isValidRookMove(board *model.Board, from model.Square, to model.Square) bool {
	piece := board.Board[from.Rank][from.File]

	// Is valid destination
	if isWhite(piece) {
		if isWhite(board.Board[to.Rank][to.File]) {
			return false
		}
	} else {
		if isBlack(board.Board[to.Rank][to.File]) {
			return false
		}
	}

	// Obey rook move rules
	if from.Rank == to.Rank {
		if from.File < to.File {
			for i := from.File + 1; i < to.File; i++ {
				if board.Board[from.Rank][i] != model.Piece(' ') {
					return false
				}
			}
		} else {
			for i := from.File - 1; i > to.File; i-- {
				if board.Board[from.Rank][i] != model.Piece(' ') {
					return false
				}
			}
		}
	} else if from.File == to.File {
		if from.Rank < to.Rank {
			for i := from.Rank + 1; i < to.Rank; i++ {
				if board.Board[i][from.File] != model.Piece(' ') {
					return false
				}
			}
		} else {
			for i := from.Rank - 1; i > to.Rank; i-- {
				if board.Board[i][from.File] != model.Piece(' ') {
					return false
				}
			}
		}
	} else {
		return false
	}

	return true
}

func isValidKnightMove(board *model.Board, from model.Square, to model.Square) bool {
	piece := board.Board[from.Rank][from.File]

	// Is valid destination
	if isWhite(piece) {
		if isWhite(board.Board[to.Rank][to.File]) {
			return false
		}
	} else {
		if isBlack(board.Board[to.Rank][to.File]) {
			return false
		}
	}

	// Obey knight move rules
	if from.Rank == to.Rank+2 || from.Rank == to.Rank-2 {
		if from.File == to.File+1 || from.File == to.File-1 {
			return true
		}
	} else if from.Rank == to.Rank+1 || from.Rank == to.Rank-1 {
		if from.File == to.File+2 || from.File == to.File-2 {
			return true
		}
	}

	return false
}

func isValidBishopMove(board *model.Board, from model.Square, to model.Square) bool {
	piece := board.Board[from.Rank][from.File]

	// Is valid destination
	if isWhite(piece) {
		if isWhite(board.Board[to.Rank][to.File]) {
			return false
		}
	} else {
		if isBlack(board.Board[to.Rank][to.File]) {
			return false
		}
	}

	fileDelta := abs(to.File - from.File)
	rankDelta := abs(to.Rank - from.Rank)

	if fileDelta != rankDelta {
		return false
	}

	// Obey bishop move rules
	if from.Rank < to.Rank {
		if from.File < to.File {
			for i := 1; i < to.Rank-from.Rank; i++ {
				if board.Board[from.Rank+i][from.File+i] != model.Piece(' ') {
					return false
				}
			}
		} else {
			for i := 1; i < to.Rank-from.Rank; i++ {
				if board.Board[from.Rank+i][from.File-i] != model.Piece(' ') {
					return false
				}
			}
		}
	} else {
		if from.File < to.File {
			for i := 1; i < from.Rank-to.Rank; i++ {
				if board.Board[from.Rank-i][from.File+i] != model.Piece(' ') {
					return false
				}
			}
		} else {
			for i := 1; i < from.Rank-to.Rank; i++ {
				if board.Board[from.Rank-i][from.File-i] != model.Piece(' ') {
					return false
				}
			}
		}
	}

	return true
}

func isValidQueenMove(board *model.Board, from model.Square, to model.Square) bool {
	return isValidBishopMove(board, from, to) || isValidRookMove(board, from, to)
}

func isValidKingMove(board *model.Board, from model.Square, to model.Square) bool {
	piece := board.Board[from.Rank][from.File]

	// Is valid destination
	if isWhite(piece) {
		if isWhite(board.Board[to.Rank][to.File]) {
			return false
		}
	} else {
		if isBlack(board.Board[to.Rank][to.File]) {
			return false
		}
	}

	// Obey king move rules
	if abs(from.Rank-to.Rank) > 1 || abs(from.File-to.File) > 1 {
		return false
	}

	return true
}

func isValidMove(board *model.Board, from model.Square, to model.Square) bool {
	piece := board.Board[from.Rank][from.File]
	if piece == model.Piece(' ') {
		return false
	}

	if from == to {
		return false
	}

	if (isWhite(piece) && board.ActiveColor == "b") || (isBlack(piece) && board.ActiveColor == "w") {
		return false
	}

	pieceRune := rune(piece)
	return pieceValidationMap[pieceRune](board, from, to)
}
