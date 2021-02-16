function fillArray(length) {
  const arr = [];

  for (let i = 0; i < length; i++) {
    arr.push(Math.floor(Math.random() * 10));
  }

  return arr;
}

function findMaxXOR(arr) {
    let max = -Infinity;
    let elements = [];

    for (let i = 0; i < arr.length; i++) {  
        for (let j = 0; j < arr.length; j++) {
            const res = arr[i] ^ arr[j];

            if (res > max) {
                max = res;
                elements = [ arr[i], arr[j] ];
            }
        }
    }

    console.log(max);
    return elements;
}

const arr = [ 25, 10 , 2, 8, 5, 3 ];
console.log(findMaxXOR(arr));