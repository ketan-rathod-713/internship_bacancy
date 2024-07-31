console.log("comparisons in javascript")

console.log(5 > null) // true
console.log(0 == null) // false
console.log(0 >= null)  // true
console.log(undefined == null) // sticky couple, true
console.log(0 >= undefined) // false

console.log("02" > 1)  // true
console.log(1 > "2")  // false

// == and <=, >= ka kam karne ka tarika thoda alag he.
// undefined ke sath kuch bhi karne se false value hi aayegi

console.log(undefined > -3)
console.log(undefined < 1)
console.log(undefined == 0)

// above 3 will returns false

// Strict equality
// Checks data types as well as value

console.log(5 === "5")  // false
console.log(5 == "5") // true