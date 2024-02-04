package main

import (
  "chess-engine/model"
  "fmt"
  "net/http"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
  html := ""
  html += makeHTMLBoard(false, globalBoard)
  html += makeHTMLBoard(true, globalBoard)

  _, err := fmt.Fprintf(w, "%s\n", html)
  if err != nil {
    fmt.Println("Error in writing response")
  }
}

func validMoveHandler(w http.ResponseWriter, r *http.Request) {
  from := r.URL.Query().Get("from")

  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      toSquare := model.Square{
        File: file,
        Rank: rank,
      }
      if isValidMove(&globalBoard, parseSquare(from), toSquare) {
        _, err := fmt.Fprintf(w, "%c%c\n", 'a'+file, '8'-rank)
        if err != nil {
          fmt.Println("Error in writing response")
        }
      }
    }
  }
}

func currentFENHandler(w http.ResponseWriter, r *http.Request) {
  _, err := fmt.Fprintf(w, "%s\n", boardToFEN(globalBoard))
  if err != nil {
    fmt.Println("Error in writing response")
  }
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
  fen := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
  globalBoard = FENToBoard(fen)
  globalBoard.MoveList = []string{}
}

func moveListHandler(w http.ResponseWriter, r *http.Request) {
  for n, move := range globalBoard.MoveList {
    _, err := fmt.Fprintf(w, "<div class='move-list-item'>%d: %s</div>", n+1, move)
    if err != nil {
      fmt.Println("Error in writing response")
    }
  }
}

func computerMoveHandler(w http.ResponseWriter, r *http.Request) {
  computeMinimaxABPruningMove(&globalBoard)
}

func movePieceHandler(w http.ResponseWriter, r *http.Request) {
  from := r.URL.Query().Get("from")
  to := r.URL.Query().Get("to")
  _, err := fmt.Fprintf(w, "Moving piece from %s to %s\n", from, to)
  if err != nil {
    fmt.Println("Error in writing response")
  }

  ret := movePiece(&globalBoard, from, to)
  if ret {
    _, err := fmt.Fprintf(w, "Move successful\n")
    if err != nil {
      fmt.Println("Error in writing response")
    }
  } else {
    _, err := fmt.Fprintf(w, "Move failed\n")
    if err != nil {
      fmt.Println("Error in writing response")
    }
  }

  if globalBoard.ActiveColor == "b" {
    if globalBoard.BlackAI != nil {
      globalBoard.BlackAI(&globalBoard)
    }
  } else {
    if globalBoard.WhiteAI != nil {
      globalBoard.WhiteAI(&globalBoard)
    }
  }
}

var pieceValueMap = map[rune]int{
  'P': 1,
  'p': 1,
  'R': 5,
  'r': 5,
  'N': 3,
  'n': 3,
  'B': 3,
  'b': 3,
  'Q': 9,
  'q': 9,
  'K': 0,
  'k': 0,
}

func pieceValuesHandler(w http.ResponseWriter, r *http.Request) {
  sumWhite := 0
  sumBlack := 0

  for rank := 0; rank < 8; rank++ {
    for file := 0; file < 8; file++ {
      piece := globalBoard.Board[rank][file]
      if isWhite(piece) {
        sumWhite += pieceValueMap[rune(piece)]
      }
      if isBlack(piece) {
        sumBlack += pieceValueMap[rune(piece)]
      }
    }
  }

  _, err := fmt.Fprintf(w, "White: %d\n", sumWhite)
  if err != nil {
    fmt.Println("Error in writing response")
  }

  _, err = fmt.Fprintf(w, "Black: %d\n", sumBlack)
  if err != nil {
    fmt.Println("Error in writing response")
  }
}

func serve() {
  http.HandleFunc("/board", boardHandler)
  http.HandleFunc("/move", movePieceHandler)
  http.HandleFunc("/valid-moves", validMoveHandler)
  http.HandleFunc("/fen", currentFENHandler)
  http.HandleFunc("/reset", resetHandler)
  http.HandleFunc("/move-list", moveListHandler)
  http.HandleFunc("/computer-move", computerMoveHandler)
  http.HandleFunc("/piece-values", pieceValuesHandler)

  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  fmt.Printf("Starting server at port 8080\n")
  if err := http.ListenAndServe(":8080", nil); err != nil {
    fmt.Printf("Error: %s\n", err)
  }
}
