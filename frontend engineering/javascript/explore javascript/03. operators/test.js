// Task 1: Guess the output of below expressions

// "" + 1 + 0 -> "10"
// "" - 1 + 0 -> -1
// true + false -> 1
// 6 / "3" -> 2
// "2" * "3" -> 6
// 4 + 5 + "px" -> "9px"
// "$" + 4 + 5 -> "$45"
// "4" - 2 -> 2
// "4px" - 2 -> NaN
// "  -9  " + 5 -> -4  // Wrong answer, should be "   -9   5"
// "  -9  " - 5  -> -14
// null + 1 -> 1 // null is 0 here
// undefined + 1 -> 1   // Wrong answer, should be NaN
// " \t \n" - 2 -> NaN  // Wrong answer, should be -2 as that string value is 0

// Task 2: Guess the output of below expressions

let a = 2;

let x = 1 + (a *= 2);

console.log(a, x) // 4, 5

// Task 3: Fix it

// let a = prompt("First number?", 1);
// let b = prompt("Second number?", 2);

// alert(a + b); // 12

// Answer: use explicit type conversion by using Number() function