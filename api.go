package main

import (
  "fmt"
  "net/http"
)

func boardHandler(w http.ResponseWriter, r *http.Request) {
}

func validMoveHandler(w http.ResponseWriter, r *http.Request) {
}

func currentFENHandler(w http.ResponseWriter, r *http.Request) {
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
}

func moveListHandler(w http.ResponseWriter, r *http.Request) {
}

func computerMoveHandler(w http.ResponseWriter, r *http.Request) {
}

func movePieceHandler(w http.ResponseWriter, r *http.Request) {
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
}

