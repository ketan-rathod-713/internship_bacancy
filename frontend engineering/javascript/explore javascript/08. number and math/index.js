// number type and maths in javascript

const num = new Number(100)
console.log(num)

// chain operation is allowed in javascript
console.log(num.toString().length)

// for precision values
console.log(num.toFixed(2))


// carefully use precision as ye decimal ke pehle ka hi dekhta he
const otherNumber = 23.8966
console.log(otherNumber.toPrecision(3))

const otherNumber2 = 233.563
console.log(otherNumber2.toPrecision(3))

// it will output in terms of 10 power for next 4th.
const otherNumber3 = 2334.5
console.log(otherNumber3.toPrecision(3))
// TODO: toPrecision in javascript

// show numbers in different formats
const hundreds = 1000000
console.log(hundreds.toLocaleString("en-IN"))

console.log(Number.MIN_VALUE, Number.MAX_VALUE, Number.isSafeInteger(3))


// MATHS in javascript
// by default it comes with javascript

// Math object gives all the methods and constants required for mathematical operations
console.log(Math)

// get random values between min and max

const min = 10
const max = 20

// it will output values in range [10, 19]
const value = Math.floor(Math.random() * (max - min)) + min
console.log(value)