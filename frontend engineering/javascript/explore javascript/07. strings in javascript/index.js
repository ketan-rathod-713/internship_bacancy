console.log("strings in javascript")

console.log(1 + 1 + "2")
console.log(2 + "2" + 2)
console.log(1 + 1 + "2" + 2)

// ECMA stardard defines ki konse operation se kisko preference milna chahiye for automatic type conversions

console.log(+true)

console.log(false)

console.log(-false)

console.log(typeof -false)

// Negative zero ha ha
let value = -false
console.log(typeof value, value)

// readability is the most important thing in programming

// Strings in javascript
const name = "hitesh"
const age = 22

// old way
// console.log(name + age + "value")

// using backticks
console.log(`name:${name} age:${age} value`)

// another way of declaring string
const gameName = new String("hello i am string")
console.log(gameName)

// key value pair me store hota he // 0 1 2 and char h e l
console.log(gameName[0])

// empty object
console.log(gameName.__proto__)

// get length of string 
console.log(gameName.length)

// prototype ke andar ye method he fir bhi me yaha se access kar sakta hu sari methods defined inside the prototype
console.log(gameName.toUpperCase())

// konsi position pe konsa character he
console.log(gameName.charAt(1))

// indexOf -> position of a character
console.log(gameName.indexOf("l"))

// TODO: list down all the methods of the string in javascript

// slice, last vali position include nahi hogi
const newString = gameName.substring(0, 4)
console.log(newString)

const anotherString = gameName.slice(0, 4)
console.log(anotherString)

// slice is same as substring but in slice we can also use negative index
console.log(gameName.slice(-5, -1))

const newStringOne = "  good    "
// how to trim it

console.log(newStringOne)
// only works on line space and line end characters
console.log(newStringOne.trim())

// browser spaces ko nahi samjta he so 
const url = "https://www.google.com/search?q=javascript%20hitesh"

console.log(url.replace("%20", " "))

// to check if string contains something
console.log(url.includes("hitesh"))

// convert string to array by using split method
const str = "hello-ketan-rathod"
console.log(str.split("-"))

