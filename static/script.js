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

    let url=`/valid-moves?from=${file}${rank}`;
    fetch(url)
      .then(response => response.text())
      .then(data => {
        moves = data.split('\n');
        for (let square of moves) {
          if (square == '') {
            continue;
          }
          let squareElements = document.querySelectorAll(`#${square}`);
          for (let squareElement of squareElements) {
            squareElement.classList.add('valid');
          }
        }
      });

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
  let urlParams = new URLSearchParams(window.location.search);
  let flipped = urlParams.get('flipped');
  if (flipped == 'true') {
    document.querySelector('#white-pov').style.display = '';
  } else {
    document.querySelector('#black-pov').style.display = '';
  }

  let moveListItems = document.querySelectorAll('.move-list-item');
  for (let moveListItem of moveListItems) {
    moveListItem.addEventListener('mouseenter', function() {
      let move = moveListItem.innerHTML.split(' ')[1];
      highlight_move(move);
    });

    moveListItem.addEventListener('mouseleave', function() {
      let move = moveListItem.innerHTML.split(' ')[1];
      highlight_move(move, false);
    });
  }

  let squares = document.querySelectorAll('.square');
  for (let square of squares) {
    square.addEventListener('mousedown', function(e) {
      if (e.button == 2) {
        origin = document.elementFromPoint(e.clientX, e.clientY);
        return;
      }
    });
    square.addEventListener('contextmenu', function(e) {
      e.preventDefault();
    });
  }
}

function createArrow() {
  let arrow = document.createElement('div');
  arrow.classList.add('arrow');
  arrow.style.transformOrigin = '0 50%';
  arrow.style.transform = 'rotate(45deg)';
  arrow.style.zIndex = '1000';
  arrow.id = 'arrow';
  document.body.appendChild(arrow);
  return arrow;
}

createArrow();

document.addEventListener('mousemove', function(e) {
  if (origin == null) {
    return;
  }

  let square = document.elementFromPoint(e.clientX, e.clientY);
  if (square == null) {
    return;
  }

  if (square.classList.contains('square')) {
    // draw arrow
    let originRect = origin.getBoundingClientRect();
    let squareRect = square.getBoundingClientRect();
    let originX = originRect.x + originRect.width / 2;
    let originY = originRect.y + originRect.height / 2;
    let squareX = squareRect.x + squareRect.width / 2;
    let squareY = squareRect.y + squareRect.height / 2;
    let angle = Math.atan2(squareY - originY, squareX - originX);
    let distance = Math.sqrt(Math.pow(squareX - originX, 2) + Math.pow(squareY - originY, 2));
    let arrow = document.querySelector('#arrow');
    arrow.style.display = '';
    arrow.style.left = `${originX}px`;
    arrow.style.top = `${originY}px`;
    arrow.style.width = `${distance}px`;
    arrow.style.transform = `rotate(${angle}rad)`;
  } else {
    let arrow = document.querySelector('#arrow');
    arrow.style.display = 'none';
  }
});

document.addEventListener('mouseup', function() {
  origin = null;
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
