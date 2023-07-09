let picked=null;
let pickedPieces=null;
let currentFen=null;
let moves=null;
let origin=null;

fetch('/piece-values')
  .then(response => response.text())
  .then(data => {
    document.querySelector('#position-values').innerHTML = data;
  });

fetch('/fen')
  .then(response => response.text())
  .then(data => {
    currentFen = data;
    document.querySelector('#fen').innerHTML = data;
    let player = null;
    if (data.split(' ')[1] == 'w') {
      player = 'White to move.';
    } else {
      player = 'Black to move.';
    }
    document.querySelector('#current-turn').innerHTML = `<h2>${player}</h2>`;
  });

function highlight_move(move, highlight=true) {
  let squares = move.split('-');
  let from = squares[0];
  let to = squares[1];
  let fromElements = document.querySelectorAll(`#${from}`);
  for (let fromElement of fromElements) {
    if (highlight) {
      fromElement.classList.add('last-move-from-temp');
    } else {
      fromElement.classList.remove('last-move-from-temp');
    }
  }
  let toElements = document.querySelectorAll(`#${to}`);
  for (let toElement of toElements) {
    if (highlight) {
      toElement.classList.add('last-move-to-temp');
    } else {
      toElement.classList.remove('last-move-to-temp');
    }
  }
}

function pickSquare(file, rank) {
}

function reset() {
}

function flip() {
}

function computer_move() {
}

window.onload = function() {
}

function createArrow() {
}

createArrow();

document.addEventListener('mousemove', function(e) {
});

document.addEventListener('mouseup', function() {
});

setInterval(function() {
  let url='/fen';
  fetch(url)
    .then(response => response.text())
    .then(data => {
      if (currentFen != data) {
        window.location.href = window.location.href;
      }
    });
}, 1000);
