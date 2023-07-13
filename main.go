package main

import (
  "chess-engine/api"
  "fmt"
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

func makeHTMLBoard(flipped bool, board Board) string {

  if (!flipped) {
    counter := 1
    html := "<div class='board' style='display:none;' id='black-pov'>"
    html += "<div class='row'>"
    html += "<span class='square'></span>"
    for file := 0; file < 8; file++ {
      html += fmt.Sprintf("<span class='square'>%c</span>", 'a' + file)
    }
    html += "</div>"
    for rank := 0; rank < 8; rank++ {
      html += "<div class='row'>"
      html += fmt.Sprintf("<span class='square'>%c</span>", '8' - rank)
      for file := 0; file < 8; file++ {
        piece := pieceToEmoji(board.board[rank][file])
        letterFile := 'a' + file

        foo := ""
        if (board.lastMove[0] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
          foo += "last-move-from "
        }
        if (board.lastMove[1] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
          foo += "last-move-to "
        }
        if (counter % 2 == 0) {
          foo += "light"
        } else {
          foo += "dark"
        }

        if piece == " " {
          html += fmt.Sprintf("<span class='square %s empty' onclick='pickSquare(\"%c\", %c)' id='%c%c'></span>", foo, letterFile, '8' - rank, letterFile, '8' - rank)
        } else {
          html += fmt.Sprintf("<span class='square %s' onclick='pickSquare(\"%c\", %c)' id='%c%c'>%s</span>", foo, letterFile, '8' - rank, letterFile, '8' - rank, piece)
        }
        counter++
      }
      html += "</div>"
      counter++
    }
    html += "</div>"

    return html
  } else {
    counter := 1
    html := "<div class='board' style='display:none;' id='white-pov'>"
    html += "<div class='row'>"
    html += "<span class='square'></span>"
    for file := 7; file >= 0; file-- {
      html += fmt.Sprintf("<span class='square'>%c</span>", 'a' + file)
    }
    html += "</div>"
    for rank := 7; rank >= 0; rank-- {
      html += "<div class='row'>"
      html += fmt.Sprintf("<span class='square'>%c</span>", '8' - rank)
      for file := 7; file >= 0; file-- {
        piece := pieceToEmoji(board.board[rank][file])
        letterFile := 'a' + file

        foo := ""
        if (board.lastMove[0] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
          foo += "last-move-from "
        }
        if (board.lastMove[1] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
          foo += "last-move-to "
        }
        if (counter % 2 == 0) {
          foo += "light"
        } else {
          foo += "dark"
        }

        if piece == " " {
          html += fmt.Sprintf("<span class='square %s empty' onclick='pickSquare(\"%c\", %c)' id='%c%c'></span>", foo, letterFile, '8' - rank, letterFile, '8' - rank)
        } else {
          html += fmt.Sprintf("<span class='square %s' onclick='pickSquare(\"%c\", %c)' id='%c%c'>%s</span>", foo, letterFile, '8' - rank, letterFile, '8' - rank, piece)
        }
        counter++
      }
      html += "</div>"
      counter++
    }
    html += "</div>"

    return html
  }
}

func main() {
  api.Foo();

  resetHandler(nil, nil)

  globalBoard.blackAI = computeMinimaxABPruningMove

  serve();
}
