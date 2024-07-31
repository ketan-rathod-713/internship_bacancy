let value;

console.log(value) // outputs undefined

value = 10n; // BigInt
console.log(typeof value) // bigint

// how to compare types
if (typeof value == 'bigint'){
    console.log("value is a BigInt")
}

// invalid operation hence NaN
console.log("i am a string"/ 20)

// dividing by zero means Infinity
console.log(1/0)

// there are total 2 types of data types
// 1. premitive data types
// 2. reference data types ( no primitive data types)

// kis tarah se data ko memory me rakha jata he uske hisab se data types k 2 types he.

// primitive data types are call by value
// 7 types : string, number, boolean, null, undefined, symbol, and bigint
// kisi bhi value ko unique banane ke liye symbol use kiya jata he.


// reference data types are call by reference
// Array, Object, Functions

// Javascript is dynamically typed language

const a = 100
const b = 10.20

const isLoggedIn = true
const outsideTemp = null 
let userEmail; // undefined

console.log(typeof a)

// symbol
const id = Symbol('123')
const anotherId = Symbol('123')

// in dono ki value same nahi he
// value same he inside se, but jo symbol se jo return value mila he vo same nahi he
console.log(id === anotherId)

console.log("value of id is", id, " value of anotherId is", anotherId)

const heroes = ("shaktiman", "nagraj", "doga")
let myObj = {
    name: "hitesh",
    age: 22,
}

console.log(heroes, myObj, typeof heroes, typeof myObj)

// function
// function ko ek variable me bhi store kar sakte he

// to know data type of any thing
// typeof operator is used to find the type of a variable or an expression.

// Important

console.log(typeof null) // object // null is a object
console.log(typeof undefined) // undefined
console.log(typeof null / undefined) // NaN
console.log(typeof NaN) // number
console.log(typeof true) // boolean

function myFunction(){
    console.log("I am a function")
}

console.log(typeof myFunction) // function

myFunction()