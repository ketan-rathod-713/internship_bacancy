console.log("comparisons in javascript")

// Comparisons with null
console.log(5 > null) // true
console.log(0 == null) // false
console.log(0 >= null)  // true

// STICKY COUPLE
console.log(undefined == null) // sticky couple, true

// Comparisons with string
console.log("02" > 1)  // true
console.log(1 > "2")  // false

// == and <=, >= ka kam karne ka tarika thoda alag he.
// undefined ke sath kuch bhi karne se false value hi aayegi

console.log(undefined > -3) // false
console.log(undefined < 1)  // false
console.log(undefined == 0)  // false

// above 3 will returns false

// Strict equality
// Checks data types as well as value

console.log(5 === "5")  // false
console.log(5 == "5") // true