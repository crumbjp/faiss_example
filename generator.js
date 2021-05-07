'use strict';
const _ = require('lodash');
const normalize = (vector) => {
  let l = Math.pow(_.reduce(vector, (r, v) => r + v*v, 0), 0.5);
  return vector.map(v => v/l);
};

for(let a = 1; a <= 10000; a++) {
  let v = [
    Math.pow(a, 0.1),   // increase
    Math.cos(a/1000),   // loop
    Math.pow(a, 0.09),  // increase
    Math.pow(1/a, 0.1), // decrease
    Math.pow(a, 0.08),  // increase
    Math.sin(a/1000),   // loop
    Math.cos(a/2000),   // loop
    Math.pow(1/a, 0.09),// decrease
    Math.pow(a, 0.06),  // increase
    Math.sin(a/2000),   // loop
  ];
  if(process.argv[2]) {
    for(let i in v) {
      v[i] = v[i] * (1.0 + Math.random()/100);
    }
  }
  let n = normalize(v);
  console.log(`${n[0]},${n[1]},${n[2]},${n[3]},${n[4]},${n[5]},${n[6]},${n[7]},${n[8]},${n[9]}`);
}
