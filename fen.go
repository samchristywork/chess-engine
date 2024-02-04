package main

import (
	"chess-engine/model"
	"fmt"
	"strconv"
	"strings"
)

func FENToBoard(fen string) model.Board {
	board := model.Board{}
	for rank := range board.Board {
		for file := range board.Board[rank] {
			board.Board[rank][file] = model.Piece(' ')
		}
	}

	fenParts := strings.Split(fen, " ")

	boardPart := fenParts[0]
	activeColor := fenParts[1]
	castling := fenParts[2]
	enPassant := fenParts[3]
	halfMoveClock := fenParts[4]
	fullMoveNumber := fenParts[5]

	board.ActiveColor = activeColor
	board.Castling = castling
	board.EnPassant = enPassant
	board.HalfMoveClock, _ = strconv.Atoi(halfMoveClock)
	board.FullMoveNumber, _ = strconv.Atoi(fullMoveNumber)

	ranks := strings.Split(boardPart, "/")
	for rank, fenRank := range ranks {
		file := 0
		for _, rune := range fenRank {
			if rune >= '1' && rune <= '8' {
				file += int(rune - '0')
			} else {
				board.Board[7-rank][file] = model.Piece(rune)
				file++
			}
		}
	}

	return board
}

func boardToFEN(board model.Board) string {
	fen := ""
	for rank := 7; rank >= 0; rank-- {
		emptySquares := 0
		for file := 0; file < 8; file++ {
			piece := board.Board[rank][file]
			if piece == model.Piece(' ') {
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

	fen += fmt.Sprintf(" %s %s %s %d %d",
		board.ActiveColor,
		board.Castling,
		board.EnPassant,
		board.HalfMoveClock,
		board.FullMoveNumber,
	)

	return fen
}
