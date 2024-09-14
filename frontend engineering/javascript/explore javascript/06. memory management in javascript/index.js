console.log("memory")

// There are 2 types of memory in javascript
// stack and heap

// 

let myYoutubeName = "ketan"
let anotherName = "something"

// actually isme ppass by value hota he. simply copying of value.
// hence dono ek dusare ko connected nahi he.
anotherName = myYoutubeName

myYoutubeName = "wow good"

console.log(myYoutubeName, anotherName)

// sara non primitive data type heap me rakha jata he
// unko reference kiya jata he.

let myObj = {
    name: "hitesh",
    age: 22,
}

let anotherObj = myObj

myObj.name = "hitesh kumar"

// check if both are same or not
console.log(myObj, anotherObj)

// both of them will be same, because they are referencing the same object in the heap memory.
// hence dono ke andar value change hogi.

// jitne bhi primitive variables he vo stack me hi jate he.
