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
  console.log(`Picked ${file}${rank}`);
  if (pickedPieces == null) {
    picked = [file, rank];
    pickedPieces = document.querySelectorAll(`#${file}${rank}`);
    for (let piece of pickedPieces) {
      piece.classList.add('selected');
    }
  } else {
    if (file == picked[0] && rank == picked[1]) {
      for (let piece of pickedPieces) {
        piece.classList.remove('selected');
      }

      for (let move of moves) {
        if (move == '') {
          continue;
        }
        let squareElements = document.querySelectorAll(`#${move}`);
        for (let squareElement of squareElements) {
          squareElement.classList.remove('valid');
        }
      }

      pickedPieces = null;
      picked = null;
      return;
    } else {
      let from = `${picked[0]}${picked[1]}`;
      let to = `${file}${rank}`;
      console.log(`Moved from ${from} to ${to}`);
      let url=`/move?from=${from}&to=${to}`;
      fetch(url)
        .then(_ => {
          window.location.href = window.location.href;
        })
    }
  }
}

function reset() {
  if (!confirm('Are you sure you want to reset the board?')) {
    return;
  }

  let url='/reset';
  fetch(url)
    .then(_ => {
      window.location.href = window.location.href;
    })
}

function flip() {
  let urlParams = new URLSearchParams(window.location.search);
  let flipped = urlParams.get('flipped');
  if (flipped == 'true') {
    window.location.href = "/";
  } else {
    window.location.href = "/?flipped=true";
  }
}

function computer_move() {
  let url='/computer-move';
  fetch(url)
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
