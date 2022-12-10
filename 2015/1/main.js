const fs = require('fs');

const input = fs.readFileSync('input').toString().split(/\r?\n/);

const part1 = (input) => {
  let floor = 0;
  for (let char of input[0]) {
    if (char === "(") {
      floor++;
    } else {
      floor--;
    }
  }
  return floor;
};

const part2 = (input) => {
  let floor = 0;
  let i = 0;
  for (let char of input[0]) {
    if (char === "(") {
      floor++;
    } else {
      floor--;
    }
    if (floor == -1) {
      return i + 1;
    }
    i++;
  }
};


console.log("part1 result:", part1(input));
console.log("part2 result:", part2(input));