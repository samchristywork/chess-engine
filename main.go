package main

import (
  "chess-engine/model"
  "fmt"
)

var globalBoard model.Board

func makeHTMLBoard(flipped bool, board model.Board) string {

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
        piece := pieceToEmoji(board.Board[rank][file])
        letterFile := 'a' + file

        foo := ""
        if (board.LastMove[0] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
          foo += "last-move-from "
        }
        if (board.LastMove[1] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
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
        piece := pieceToEmoji(board.Board[rank][file])
        letterFile := 'a' + file

        foo := ""
        if (board.LastMove[0] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
          foo += "last-move-from "
        }
        if (board.LastMove[1] == fmt.Sprintf("%c%c", letterFile, '8' - rank)) {
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
  resetHandler(nil, nil)

  globalBoard.BlackAI = computeMinimaxABPruningMove

  serve();
}
