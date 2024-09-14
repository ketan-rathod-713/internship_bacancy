// array is an object in javascript
// arrays in javascript are resizeable
// array elements can not be accessible by using arbitrary thing, only by using index
// 0 based indexing
// IMPORTANT : copy operations in array does shallow copy. both saves the same reference data.

// array ke andar ek prototype hota and uske andar bhi ek prototype hota he

// array basic syntax

let myarr = [1, 2, 3, 4]
console.log(myarr)

let myarr2 = new Array(1, 2, 3, 4)
console.log(myarr2)
console.log(myarr[1])

// array methods

myarr.push(6)
myarr.push(7)
let returned = myarr.pop()

// sari values ko right shift karna he
// very costly operation
myarr.unshift(9)
myarr.unshift(10)
let shifted = myarr.shift()
console.log(myarr, returned, shifted)

console.log(myarr.includes(7))

console.log(myarr.indexOf(10)) // returns -1 if not exists

// type is changed to string
const newArr = myarr.join('-')
console.log(newArr)

// IMPORTANT: slice and splice method of array in javascript

// slice
console.log("Initial Array : ", myarr)
console.log("spliced result: ",myarr.slice(1, 3))
console.log("After slice : ",myarr)

// splice
console.log("spliced result: ", myarr.splice(1, 3))
console.log("After Splice: ",myarr)