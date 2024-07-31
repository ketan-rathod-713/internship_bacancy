// Section 1 : Assignment operator

console.log("operators and their precedence")

let a = 3
let b = 1
let c = 10 - (a = b + 3)

console.log("value of a is", a)
console.log("value of c is", c)

console.log("hence value of changed in above expression from 3 to", b + 3)

// Section 2 : Chaining Assignments
// chained assignments evaluaates from right to left

a = b = c = 2 + 2;

console.log( a ); // 4
console.log( b ); // 4
console.log( c ); // 4

// Section 3 : Comma Operator

console.log("Comma operator example")

a = b = 2 + 3, 3 + 9; // it will ignore the second value

console.log("value of a is", a)

// some times it is usefull in for loop when we want to initialize multiple variables at once

for (let i = 0, j = 0; i < 5; i++, j++) {
    console.log(i, j)
}
