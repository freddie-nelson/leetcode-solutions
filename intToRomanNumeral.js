const numerals = {
  1: "I",
  4: "IV",
  5: "V",
  9: "IX",
  10: "X",
  40: "XL",
  50: "L",
  90: "XC",
  100: "C",
  400: "CD",
  500: "D",
  900: "CM",
  1000: "M",
};
const values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];

function convertIntToRomanNum(inputNumber) {
  let start = 0;
  let result = "";

  while (inputNumber > 0) {
    for (let i = start; i < values.length; i++) {
      const value = values[i];

      if (inputNumber - value >= 0) {
        inputNumber -= value;
        result += numerals[value];
        start = i;
        break;
      }
    }
  }

  return result;
}

console.time();
const string = convertIntToRomanNum(20);
console.timeEnd();

console.log(string);
